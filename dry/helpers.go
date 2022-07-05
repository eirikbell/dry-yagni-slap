package dry

func union(first, second []int) []int {
	union := map[int]struct{}{}
	for _, v := range first {
		union[v] = struct{}{}
	}
	for _, v := range second {
		union[v] = struct{}{}
	}

	keys := make([]int, 0, len(union))
	for k := range union {
		keys = append(keys, k)
	}
	return keys
}

func intersect(first, second []int) []int {
	x := map[int]struct{}{}
	for _, v := range first {
		x[v] = struct{}{}
	}
	y := map[int]struct{}{}
	for _, v := range second {
		y[v] = struct{}{}
	}

	result := []int{}
	for k, _ := range x {
		if _, ok := y[k]; ok {
			result = append(result, k)
		}
	}
	return result
}

func mapKeys(in map[int]struct{}) []int {
	keys := make([]int, 0, len(in))
	for k := range in {
		keys = append(keys, k)
	}
	return keys
}

func sliceToMap(in []int) map[int]struct{} {
	result := map[int]struct{}{}
	for _, v := range in {
		result[v] = struct{}{}
	}
	return result
}

func copy(in []int) []int {
	result := []int{}
	return append(result, in...)
}
