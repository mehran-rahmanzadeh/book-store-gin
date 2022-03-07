package utils

func CheckItemInSlice(s []string, item string) bool {
	for _, val := range s {
		if val == item {
			return true
		}
	}
	return false
}
