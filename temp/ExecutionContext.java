package com.zion.backend.support.actionflow;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.node.NullNode;
import com.fasterxml.jackson.databind.node.ObjectNode;
import com.functorz.ztype.liveschema.RunCustomCode;
import com.google.common.base.Preconditions;
import com.zion.backend.support.actionflow.graphql.TableDeleteArgs;
import com.zion.backend.support.actionflow.graphql.TableInsertArgs;
import com.zion.backend.support.actionflow.graphql.TableQueryArgs;
import com.zion.backend.support.actionflow.graphql.TableUpdateArgs;
import com.zion.backend.support.graphql.GraphQLConstants;
import com.zion.backend.support.graphql.GraphQLRequestContext;
import com.zion.backend.support.graphql.datafetcher.OpenApiExecutor;
import com.zion.backend.support.graphql.error.DetectedException;
import com.zion.backend.support.graphql.permission.GraphQLPermissionSession;
import com.zion.backend.support.polyglot.PolyglotUtils;
import com.zion.backend.support.security.wechat.WechatThirdPartyPlatformService;
import com.zion.backend.support.utils.JsonUtils;
import com.zion.backend.support.utils.Utils;
import java.util.Optional;
import org.graalvm.polyglot.Context;
import org.graalvm.polyglot.Value;
import org.springframework.context.ApplicationContext;

import java.util.HashMap;
import java.util.Map;
import java.util.Set;
import java.util.regex.Pattern;
import org.springframework.data.util.Pair;

public class ExecutionContext {
  private Map<String, Value> returnValues = new HashMap<>();
  private Map<String, Object> args;
  private String uniqueId;
  private Set<String> inputNames;
  private Set<String> outputNames;
  private ActionFlowGraphQLExecutor graphQLExecutor;
  private OpenApiExecutor openApiExecutor;
  private ActionFlowService actionFlowService;
  private boolean testOnly;
  private Context context;
  private GraphQLRequestContext requestContext;
  private WechatThirdPartyPlatformService wechatThirdPartyPlatformService;

  ExecutionContext(
      RunCustomCode customCodeNode,
      Map<String, Object> args,
      Context context,
      GraphQLRequestContext requestContext,
      ApplicationContext applicationContext) {
    this.args = Preconditions.checkNotNull(args);
    this.graphQLExecutor = Utils.getBeanByTypeOrThrow(applicationContext, ActionFlowGraphQLExecutor.class);
    this.context = context;
    this.requestContext = requestContext;
    this.openApiExecutor = Utils.getBeanByTypeOrThrow(applicationContext, OpenApiExecutor.class);
    this.actionFlowService = Utils.getBeanByTypeOrThrow(applicationContext, ActionFlowService.class);
    this.uniqueId = customCodeNode.getUniqueId();
    this.inputNames = Preconditions.checkNotNull(customCodeNode.getInputArgs().keySet());
    this.outputNames = Preconditions.checkNotNull(customCodeNode.getOutputValues().keySet());
    this.testOnly = false;
    this.wechatThirdPartyPlatformService =
        Utils.getBeanByTypeOrThrow(applicationContext, WechatThirdPartyPlatformService.class);
  }

  ExecutionContext(
      Map<String, Object> args,
      String uniqueId,
      Context context,
      GraphQLRequestContext requestContext,
      ApplicationContext applicationContext) {
    this.args = Preconditions.checkNotNull(args);
    this.uniqueId = uniqueId;
    this.context = context;
    this.graphQLExecutor = Utils.getBeanByTypeOrThrow(applicationContext, ActionFlowGraphQLExecutor.class);
    this.openApiExecutor = Utils.getBeanByTypeOrThrow(applicationContext, OpenApiExecutor.class);
    this.actionFlowService = Utils.getBeanByTypeOrThrow(applicationContext, ActionFlowService.class);
    this.requestContext = requestContext;
    this.testOnly = true;
  }

  Map<String, Value> getReturnValues() {
    return returnValues;
  }

  public String getWechatMiniAppAccessToken() {
    return wechatThirdPartyPlatformService.getAccessToken();
  }

  public Object getArg(String name) {
    if (name == null
        || (!name.startsWith(GraphQLConstants.SYSTEM_FIELD_NAME_PREFIX)
            && !testOnly
            && !inputNames.contains(name))) {
      throw new UnknownValueException(name);
    }
    Object obj = args.get(name);
    return PolyglotUtils.convertMemberToJsObject(obj, context);
  }

  public void setReturn(String name, Value value) {
    if (!testOnly && !outputNames.contains(name)) {
      throw new UnknownValueException(name);
    }
    returnValues.put(name, value);
  }

  public void throwException(Value errorType, Value errorMsg) {
    String errorTypeStr = PolyglotUtils.convertPolyglotValueToObject(errorType, String.class);
    CustomThrowErrorType customThrowErrorType = CustomThrowErrorType.valueOf(errorTypeStr);
    String msg = PolyglotUtils.convertPolyglotValueToObject(errorMsg, String.class);
    throw new DetectedException(customThrowErrorType, msg);
  }

  public Object callThirdPartyApi(String operationId, Value args) {
    Map<String, Object> arguments = PolyglotUtils.convertPolyglotValueToObject(args, Map.class);
    Pair<Integer, Optional<JsonNode>> resultByHttpResponseCode = openApiExecutor.executeOpenApi(operationId, arguments);
    JsonNode data = resultByHttpResponseCode.getSecond().orElse(NullNode.getInstance());
    ObjectNode resultNode = JsonUtils.newObjectNode();
    resultNode.put("code", resultByHttpResponseCode.getFirst());
    resultNode.set("data", resultNode);
    return PolyglotUtils.convertMemberToJsObject(resultNode, context);
  }

  public Object callActionFlow(String actionFlowId, Integer versionId, Value args) {
    JsonNode arguments = PolyglotUtils.convertPolyglotValueToObject(args, JsonNode.class);
    JsonNode result = actionFlowService.executeActionFlow(actionFlowId, versionId, arguments, requestContext);
    return PolyglotUtils.convertMemberToJsObject(result, context);
  }

  public Object query(String table, Value args, Value resultFilter, Value permission) {
    String operationName = generateOperationName("query", table);
    TableQueryArgs tableQueryArgs = new TableQueryArgs(operationName, table, args, resultFilter);
    GraphQLPermissionSession session = parseSessionFromPloyglotValue(permission);
    Object object = graphQLExecutor.queryList(tableQueryArgs, session, requestContext);
    return PolyglotUtils.convertMemberToJsObject(object, context);
  }

  public int delete(String table, Value args, Value permission) {
    String operationName = generateOperationName("delete", table);
    TableDeleteArgs tableDeleteArgs = new TableDeleteArgs(operationName, table, args);
    GraphQLPermissionSession session = parseSessionFromPloyglotValue(permission);
    return graphQLExecutor.delete(tableDeleteArgs, session, requestContext);
  }

  public int update(String table, Value args, Value permission) {
    String operationName = generateOperationName("update", table);
    TableUpdateArgs tableUpdateArgs = new TableUpdateArgs(operationName, table, args);
    GraphQLPermissionSession session = parseSessionFromPloyglotValue(permission);
    return graphQLExecutor.update(tableUpdateArgs, session, requestContext);
  }

  public Object insert(String table, Value args, Value resultFilter, Value permission) {
    String operationName = generateOperationName("insert", table);
    TableInsertArgs tableInsertArgs = new TableInsertArgs(operationName, table, args, resultFilter);
    GraphQLPermissionSession session = parseSessionFromPloyglotValue(permission);
    Object returns = graphQLExecutor.insert(tableInsertArgs, session, requestContext);
    return PolyglotUtils.convertMemberToJsObject(returns, context);
  }

  public Object runGql(String operationName, String gql, Value variables, Value permission) {

    GraphQLPermissionSession session = parseSessionFromPloyglotValue(permission);
    ObjectNode node = Utils.uncheckedCast(PolyglotUtils.convertPolyglotValueToJsonNode(variables));
    Map<String, Object> variableMap =
        JsonUtils.convertFromJsonNode(node, new TypeReference<Map<String, Object>>() {});
    Object result =
        graphQLExecutor.runGql(operationName, gql, variableMap, requestContext, session);
    return PolyglotUtils.convertMemberToJsObject(result, context);
  }

  public int getSeqNextValue(String seqName, boolean createIfNotExists) {
    return graphQLExecutor.getSeqNextValue(seqName, createIfNotExists);
  }

  public void resetSeqValue(String seqName, int value) {
    graphQLExecutor.resetSeqValue(seqName, value);
  }

  public void sendEmail(
      String toAddress, String subject, String fromAlias, String textBody, String htmlBody) {
    graphQLExecutor.sendEmail(toAddress, subject, fromAlias, textBody, htmlBody);
  }

  public Object uploadMedia(String url, Value headers) {
    JsonNode headersJson = PolyglotUtils.convertPolyglotValueToJsonNode(headers);
    Map<String, String> headersToUse =
        JsonUtils.convertFromJsonNode(headersJson, new TypeReference<Map<String, String>>() {});
    return PolyglotUtils.convertMemberToJsObject(
        graphQLExecutor.uploadMediaWithUrl(url, headersToUse), context);
  }

  private GraphQLPermissionSession parseSessionFromPloyglotValue(Value permission) {
    if (!permission.hasMember(GraphQLConstants.ROLE)) {
      throw new IllegalStateException("permission should have field role");
    }
    String role =
        PolyglotUtils.convertPolyglotValueToObject(
            permission.getMember(GraphQLConstants.ROLE), String.class);
    @SuppressWarnings("unchecked")
    Map<String, Object> sessionVariable =
        PolyglotUtils.findMemberFromPolyglotValue(
                permission, GraphQLConstants.SESSION_VARIABLE, Map.class)
            .map(map -> (Map<String, Object>) map)
            .orElseGet(Map::of);
    return new GraphQLPermissionSession(role, sessionVariable);
  }

  private String generateOperationName(String operation, String tableName) {
    String uniqueIdReplaceHyphenToUnderCross = uniqueId.replaceAll(Pattern.quote("-"), "_");
    return String.format(
        "action_%s_%s_%s", uniqueIdReplaceHyphenToUnderCross, operation, tableName);
  }
}
