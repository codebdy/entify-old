package utils

func MapStringKeys(m map[string]interface{}, wapper string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, wapper+k+wapper)
	}
	return keys
}
