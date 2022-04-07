package meta

const (
	SCRIPT         string = "script"
	CLOUD_FUNCTION string = "cloudFunction"
	MICRO_SERVICE  string = "microService"

	QUERY    string = "query"
	MUTATION string = "mutation"
)

type ArgMeta struct {
	Uuid      string `json:"uuid"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	TypeUuid  string `json:"typeUuid"`
	TypeLabel string `json:"typeLabel"`
}

type MethodMeta struct {
	Uuid             string    `json:"uuid"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	TypeUuid         string    `json:"typeUuid"`
	TypeLabel        string    `json:"typeLabel"`
	Args             []ArgMeta `json:"args"`
	OperateType      string    `json:"operateType"` //Mutation or Query
	ImplementType    string    `json:"implementType"`
	MethodImplements string    `json:"methodImplements"`
	Description      string    `json:"description"`
}
