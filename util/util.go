package util

func InSlice(data []int, find int) bool {
	exists := false
	for _, item := range data {
		if item == find {
			exists = true
			break
		}
	}

	return exists
}