package domain

const (
	SCRIPT         string = "script"
	CLOUD_FUNCTION string = "cloudFunction"
	MICRO_SERVICE  string = "microService"
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
	ImplementType    string    `json:"implementType"`
	MethodImplements string    `json:"methodImplements"`
}
