package com.zion.backend.support.actionflow;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ObjectNode;
import com.functorz.ztype.eval.EvalContext;
import com.functorz.ztype.eval.EvaluatorKt;
import com.functorz.ztype.expr.ZExpr;
import com.functorz.ztype.liveschema.ActionFlow;
import com.functorz.ztype.liveschema.ActionFlowNode;
import com.functorz.ztype.liveschema.ActionFlowNodeType;
import com.functorz.ztype.liveschema.BoolExp;
import com.functorz.ztype.liveschema.BranchItem;
import com.functorz.ztype.liveschema.BranchMerge;
import com.functorz.ztype.liveschema.BranchSeparation;
import com.functorz.ztype.liveschema.QueryRecord;
import com.functorz.ztype.liveschema.RunCustomCode;
import com.functorz.ztype.liveschema.SchemaResolver;
import com.functorz.ztype.typecontext.ZTypeContext;
import com.functorz.ztype.utils.MediaEvalType;
import com.functorz.ztype.zschema.ActionFlowDefinition;
import com.functorz.ztype.zschema.BackendAvailableSchema;
import com.zion.backend.support.callback.CallbackAction;
import com.zion.backend.support.callback.action.InvokeActionFlow;
import com.zion.backend.support.graphql.GraphQLRequestContext;
import com.zion.backend.support.graphql.config.datamodel.DataModelHolder;
import com.zion.backend.support.graphql.config.datamodel.DataModelRegistry;
import com.zion.backend.support.graphql.error.DetectedException;
import com.zion.backend.support.polyglot.AbstractContext.DelegateContext;
import com.zion.backend.support.polyglot.AbstractValue;
import com.zion.backend.support.polyglot.PolyglotConstants;
import com.zion.backend.support.polyglot.PolyglotUtils;
import com.zion.backend.support.polyglot.js.JsExecutor;
import com.zion.backend.support.utils.MultiStageStopWatch;
import com.zion.backend.support.utils.Utils;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Optional;
import java.util.UUID;
import java.util.function.Function;
import java.util.stream.Collectors;
import javax.ws.rs.NotSupportedException;
import lombok.extern.slf4j.Slf4j;
import org.graalvm.polyglot.PolyglotException;
import org.graalvm.polyglot.Value;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.ApplicationContext;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@Slf4j
public class ActionFlowService implements CallbackAction.Validator {
  @Autowired
  private ObjectMapper objectMapper;
  @Autowired
  private ActionFlowRepository actionFlowRepo;
  @Autowired
  private ActionFlowGraphQLExecutor graphQLExecutor;
  @Autowired
  private ApplicationContext applicationContext;

  @Autowired
  private DataModelRegistry dataModelRegistry;

  @Autowired
  private QueryRecordExecutor queryRecordExecutor;

  @Transactional
  public JsonNode executeActionFlow(
      String actionFlowId, Integer versionId, JsonNode args, GraphQLRequestContext context) {
    Optional<ActionFlowWrapper> actionFlowWrapper =
        versionId == null
            ? actionFlowRepo.findTopByUniqueIdOrderByVersionIdDesc(UUID.fromString(actionFlowId))
            : actionFlowRepo.findByUniqueIdAndVersionId(UUID.fromString(actionFlowId), versionId);
    ActionFlow actionFlow =
        actionFlowWrapper
            .orElseThrow(
                () ->
                    new ActionFlowException(
                        ActionFlowErrorType.ACTION_FLOW_NOT_FOUND, "Action flow not found"))
            .getActionFlowContent();
    return doExecuteActionFlow(actionFlow, args, context);
  }

  @Transactional
  public JsonNode executeTestCode(
      String code,
      ObjectNode args,
      Long accountId,
      boolean updateDb,
      GraphQLRequestContext requestContext) {
    Map<String, Object> currentArgs =
        objectMapper.convertValue(
            Utils.coalesce(args, objectMapper.createObjectNode()),
            new TypeReference<Map<String, Object>>() {
            });
    if (accountId != null) {
      currentArgs.put("fzCurrentAccountId", accountId);
    }
    JsonNode result =
        JsExecutor.<JsonNode>withContext(
            context -> {
              ExecutionContext injectedContext =
                  new ExecutionContext(
                      currentArgs,
                      Utils.generateUuidStr(),
                      context.getContext(),
                      requestContext,
                      applicationContext);
              return objectMapper.valueToTree(
                  doExecution(code, currentArgs, context, injectedContext));
            });
    if (!updateDb) {
      log.info("action flow result: {}", result.toString());
      throw new RuntimeException("hack to revert db changes");
    }
    return result;
  }

  JsonNode doExecuteActionFlow(
      ActionFlow actionFlow, JsonNode args, GraphQLRequestContext context) {
    MultiStageStopWatch stopwatch =
        MultiStageStopWatch.startTiming(
            actionFlow.getDisplayName() + "-" + actionFlow.getUniqueId());
    ActionFlowDefinition actionFlowDefinition = SchemaResolver.INSTANCE.resolveActionFlow(actionFlow);
    Map<String, ActionFlowDefinition> actionFlowDefinitionByUniqueId = Map.of(actionFlowDefinition.getId(), actionFlowDefinition);
    Map<String, ActionFlowNode> flowNodesById =
        actionFlow.getAllNodes().stream()
            .collect(Collectors.toMap(ActionFlowNode::getUniqueId, Function.identity()));
    return EvalContext.Companion.assuming(new EvalContext(new ArrayList<>(), MediaEvalType.ID, actionFlowDefinition.getId()), () -> {
      DataModelHolder dataModelHolder = dataModelRegistry.getCurrentDataModelHolder();
      BackendAvailableSchema backendAvailableSchema = new BackendAvailableSchema(
          SchemaResolver.INSTANCE.resolveDataModel(dataModelHolder.getLiveSchemaDataModel()),
          new HashMap<>(),
          actionFlowDefinitionByUniqueId
      );
      return ZTypeContext.Companion.withPartialSchema(backendAvailableSchema, () -> {
        Map<String, Object> currentArgs = objectMapper.convertValue(args, new TypeReference<Map<String, Object>>() {
        });
        if (context != null && context.isAuthenticated()) {
          currentArgs.put("fzCurrentAccountId", context.getAuthenticatedAccountId());
        }
        ActionFlowNode currentNode = flowNodesById.get(actionFlow.getStartNodeId());
        stopwatch.markStageEnd("obtainArgs");
        while (currentNode != null && currentNode.getType() != ActionFlowNodeType.FLOW_END) {
          NodeExecutionResult result;
          switch (currentNode.getType()) {
            case CUSTOM_CODE -> result = executeCustomCodeNode((RunCustomCode) currentNode, currentArgs, context);
            case BRANCH_SEPARATION ->
                result = executeBranchSeparationNode((BranchSeparation) currentNode, currentArgs, flowNodesById);
            case BRANCH_MERGE -> result = executeBranchMergeNode((BranchMerge) currentNode, currentArgs);
            case QUERY_RECORD -> result = queryRecordExecutor.executeQueryRecord((QueryRecord) currentNode, currentArgs, context,
                actionFlow.getUniqueId(), backendAvailableSchema);
            default -> throw new RuntimeException("shouldn't get here");
          }
          stopwatch.markStageEnd(currentNode.getType() + "-" + currentNode.getUniqueId());
          currentNode = flowNodesById.get(result.getNextNodeId());
          currentArgs.putAll(result.getOutputValues());
        }

        JsonNode result = objectMapper.valueToTree(currentArgs);
        String timingInfo = stopwatch.markOperationEnd();
        log.info(timingInfo + ", args:" + args);
        return result;
      });
    });
  }

  private NodeExecutionResult executeCustomCodeNode(
      RunCustomCode node, Map<String, Object> args, GraphQLRequestContext requestContext) {
    String code = node.getCode();
    try {
      return JsExecutor.withContext(
          context -> {
            ExecutionContext injectedContext =
                new ExecutionContext(
                    node,
                    args,
                    context.getContext(),
                    requestContext,
                    applicationContext);
            Map<String, Object> returnValues = doExecution(code, args, context, injectedContext);
            node.getOutputValues().entrySet().stream()
                .forEach(
                    entry -> {
                      if (!entry.getValue().getNullable()
                          && !returnValues.containsKey(entry.getKey())) {
                        throw new MissingReturnValueException(entry.getKey());
                      }
                    });
            return new NodeExecutionResult(node.getAndThenNodeId(), returnValues);
          });
    } catch (PolyglotException e) {
      if (e.isHostException()) {
        Throwable wrappedException = e.asHostException();
        if (wrappedException instanceof DetectedException) {
          throw (DetectedException) wrappedException;
        }
      }
      log.error("action flow failed", e);
      throw new GenericActionFlowExcutionException(e);
    }
  }

  private NodeExecutionResult executeBranchMergeNode(BranchMerge node, Map<String, Object> args) {
    return new NodeExecutionResult(node.getAndThenNodeId(), args);
  }



  private NodeExecutionResult executeBranchSeparationNode(BranchSeparation node,
                                                          Map<String, Object> args,
                                                          Map<String, ActionFlowNode> flowNodesById) {
    switch (node.getConditionType()) {
      case MUTUAL_TOLERANCE -> throw new NotSupportedException("Tolerance branch is not supported now.");
      case MUTUAL_EXCLUSION -> {
        for (String branchItemId : node.getBranchItemIds()) {
          BranchItem branchItem = (BranchItem) flowNodesById.get(branchItemId);
          if (executeConditionals(branchItem.getCondition(), args)) {
            return new NodeExecutionResult(branchItem.getAndThenNodeId(), args);
          }
        }
      }
    }
    throw new NotSupportedException("Unknown Condition");
  }

  private boolean executeConditionals(BoolExp boolExp, Map<String, Object> inputArgs) {
    ZExpr zExpr = SchemaResolver.INSTANCE.resolveBoolExp(boolExp);
    String code = EvaluatorKt.getEvalString(zExpr, null);
    return doExecutionConditionals(code, inputArgs);
  }


  private boolean doExecutionConditionals(String code, Map<String, Object> inputArgs) {
    return JsExecutor.withContext(

        //todo: let left and right both literal first, and then change them to variable
        delegateContext -> {
          AbstractValue jsBinding = delegateContext.getBindings(JsExecutor.JS);
          jsBinding.putMember("inputArgs", inputArgs);
          Value eval = delegateContext.eval(JsExecutor.JS, code);
          return eval.getMember("data").asBoolean();
        }
    );
  }

  private Map<String, Object> doExecution(
      String code,
      Map<String, Object> args,
      DelegateContext context,
      ExecutionContext injectedContext) {
    AbstractValue jsBinding = context.getBindings(JsExecutor.JS);
    jsBinding.putMember(PolyglotConstants.INJECTED_CONTEXT_NAME, injectedContext);
    try {
      context.eval(JsExecutor.JS, code);
    } catch (RuntimeException e) {
      if(e instanceof PolyglotException){
        PolyglotException polyglotException=(PolyglotException) e;
        if(polyglotException.asHostException() instanceof DetectedException){
          throw (DetectedException) ((PolyglotException) e).asHostException();
        }
      }
      throw e;
    }
    Map<String, Object> returnValues =
        convertValuesToPlainObjects(injectedContext.getReturnValues());
    return returnValues;
  }

  private Map<String, Object> convertValuesToPlainObjects(Map<String, Value> values) {
    Map<String, JsonNode> asJson =
        values.entrySet().stream()
            .collect(
                Collectors.toMap(
                    Map.Entry::getKey,
                    entry -> PolyglotUtils.convertPolyglotValueToJsonNode(entry.getValue())));
    return objectMapper.convertValue(asJson, new TypeReference<Map<String, Object>>() {
    });
  }

  @Override
  public void validate(CallbackAction action) {
    if (!(action instanceof InvokeActionFlow)) {
      return;
    }
    InvokeActionFlow actionFlow = Utils.uncheckedCast(action);
    boolean exist =
        actionFlowRepo.existsByUniqueId(UUID.fromString(actionFlow.getActionFlowUniqueId()));
    if (!exist) {
      String msg =
          String.format(
              "action flow with uniqueId: %s and version: %s",
              actionFlow.getActionFlowUniqueId(), actionFlow.getVersion());
      throw new IllegalArgumentException(msg);
    }
  }

  public static String getGqlOpNameByNodeId(String nodeId) {
    return String.format(ActionFlowConstants.GQL_OPERATION_NAME_FORMAT, nodeId.replace("-", "_"));
  }
}
