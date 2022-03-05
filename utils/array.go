package utils

func StringFilter(arr []string, f func(value string) bool) []string {
	positive := []string{}

	for i := range arr {
		if f(arr[i]) {
			positive = append(positive, arr[i])
		}
	}

	return positive
}
