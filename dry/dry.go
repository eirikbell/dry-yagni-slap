package dry

import "sort"

// dont mutate
var listA = []int{3, 4, 1, 6, 8, 7, 1, 1, 5}
var listB = []int{9, 9, 7, 7, 5, 5, 3, 3}

func sortedUniqueUnion() []int {
	a := []int{}
	for _, val := range listA {
		a = append(a, val)
	}
	b := []int{}
	for _, val := range listB {
		b = append(b, val)
	}

	uniqueA := []int{}
	previous := 0
	for i, current := range a {
		if i > 0 && current == previous {
			continue
		}
		uniqueA = append(uniqueA, current)
		previous = current
	}

	uniqueB := []int{}
	previous = 0
	for i, current := range b {
		if i > 0 && current == previous {
			continue
		}
		uniqueB = append(uniqueB, current)
		previous = current
	}

	union := map[int]struct{}{}
	for _, valA := range uniqueA {
		union[valA] = struct{}{}
	}
	for _, valB := range uniqueB {
		union[valB] = struct{}{}
	}

	result := []int{}
	for k, _ := range union {
		result = append(result, k)
	}
	sort.Ints(result)
	return result
}

func sortedDescendingUniqueIntersect() []int {
	a := []int{}
	for _, val := range listA {
		a = append(a, val)
	}
	b := []int{}
	for _, val := range listB {
		b = append(b, val)
	}

	uniqueA := []int{}
	previous := 0
	for i, current := range a {
		if i > 0 && current == previous {
			continue
		}
		uniqueA = append(uniqueA, current)
		previous = current
	}

	uniqueB := []int{}
	previous = 0
	for i, current := range b {
		if i > 0 && current == previous {
			continue
		}
		uniqueB = append(uniqueB, current)
		previous = current
	}

	mapA := map[int]struct{}{}
	for _, valA := range uniqueA {
		mapA[valA] = struct{}{}
	}
	mapB := map[int]struct{}{}
	for _, valB := range uniqueB {
		mapB[valB] = struct{}{}
	}

	intersect := []int{}
	for k, _ := range mapA {
		if _, ok := mapB[k]; ok {
			intersect = append(intersect, k)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(intersect)))
	return intersect
}
