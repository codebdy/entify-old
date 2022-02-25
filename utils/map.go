package utils

func MapStringKeys(m map[string]interface{}, wapper string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, wapper+k+wapper)
	}
	return keys
}

func MapValues(m map[string]interface{}, wapper string) []interface{} {
	values := make([]interface{}, 0, len(m))
	for k := range m {
		values = append(values, m[k])
	}
	return values
}
