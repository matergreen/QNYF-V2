package utils

// 切片去重 去空
func UniqueSlice(s []string) []string {
	result := make([]string, 0)
	tmpMap := make(map[string]bool, len(s))
	for _, tmp := range s {
		if tmp == "" {
			continue
		}
		if _, ok := tmpMap[tmp]; !ok {
			result = append(result, tmp)
			tmpMap[tmp] = true
		}
	}
	return result
}
