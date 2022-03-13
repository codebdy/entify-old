package meta

const (
	COLUMN_ID           string = "ID"
	COLUMN_INT          string = "Int"
	COLUMN_FLOAT        string = "Float"
	COLUMN_BOOLEAN      string = "Boolean"
	COLUMN_STRING       string = "String"
	COLUMN_DATE         string = "Date"
	COLUMN_SIMPLE_JSON  string = "SimpleJson"
	COLUMN_SIMPLE_ARRAY string = "simpleArray"
	COLUMN_JSON_ARRAY   string = "JsonArray"
	COLUMN_ENUM         string = "Enum"
)

type Column struct {
	Uuid          string `json:"uuid"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Primary       bool   `json:"primary"`
	Generated     bool   `json:"generated"`
	Nullable      bool   `json:"nullable"`
	Unique        bool   `json:"unique"`
	Index         bool   `json:"index"`
	CreateDate    bool   `json:"createDate"`
	UpdateDate    bool   `json:"updateDate"`
	DeleteDate    bool   `json:"deleteDate"`
	Select        bool   `json:"select"`
	Length        int    `json:"length"`
	FloatM        int    `json:"floatM"` //M digits in total
	FloatD        int    `json:"floatD"` //D digits may be after the decimal point
	Unsigned      bool   `json:"unsigned"`
	TypeEnityUuid string `json:"typeEnityUuid"`
	Description   string `json:"description"`
}
