package spy

func spySolution(clothes [][]string) int {
	answer := 0
	var kind map[string]int
	kind = make(map[string]int)

	for _, c := range clothes {
		k := c[1]
		_, exists := kind[k]

		if !exists {
			kind[k] = 0
		}

		kind[k]++
	}

	result := 1

	for _, v := range kind {
		result *= v + 1
	}

	answer = result - 1

	return answer
}
