package dialect

const (
	MySQL = "mysql"
)

type SQLBuilder interface {
	BuildBoolExp(where map[string]interface{}) (string, []interface{})
	BuildFieldExp(fieldName string, fieldArgs map[string]interface{}) (string, []interface{})
}

func GetSQLBuilder() SQLBuilder {
	var builder MySQLBuilder
	return &builder
}
