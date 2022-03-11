package dialect

type MySQLBuilder struct {
}

func (*MySQLBuilder) BuildBoolExp(where map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	return "", params
}
