package schema

const (
	COLUMN_NUMBER       string = "Number"
	COLUMN_BOOLEAN      string = "Boolean"
	COLUMN_STRING       string = "String"
	COLUMN_TEXT         string = "Text"
	COLUMN_MEDIUM_TEXT  string = "MediumText"
	COLUMN_LONG_TEXT    string = "LongText"
	COLUMN_DATE         string = "Date"
	COLUMN_SIMPLE_JSON  string = "SimpleJson"
	COLUMN_SIMPLE_ARRAY string = "simpleArray"
	COLUMN_JSON_ARRAY   string = "JsonArray"
	COLUMN_ENUM         string = "Enum"
)

type ColumnMeta struct {
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
	TypeEnityUuid string `json:"typeEnityUuid"`
	Description   string `json:"description"`
}
