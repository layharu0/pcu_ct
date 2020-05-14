package album

func albumSolution(genres []string, plays []int) []int {
	answer := make([]int, 0)

	totalPlay := make(map[string]int)
	musicList := make(map[string][]int)
	for i := 0; i < len(genres); i++ {
		g := genres[i]

		_, exists := totalPlay[g]
		if !exists {
			totalPlay[g] = 0
			musicList[g] = make([]int, 0)
		}

		totalPlay[g] += plays[i]
		musicList[g] = append(musicList[g], i)
	}

	genreSort := make([]string, 0)
	for name, play := range totalPlay {
		if len(genreSort) < 1 {
			genreSort = append(genreSort, name)
		} else {
			for i := 0; i < len(genreSort); i++ {
				g := genreSort[i]
				if play > totalPlay[g] {
					tmp := genreSort
					genreSort = append(make([]string, 0), name)
					genreSort = append(genreSort, tmp...)
					break
				}

				if i == len(genreSort)-1 {
					genreSort = append(genreSort, name)
				}
			}
		}
	}

	for _, g := range genreSort {
		first := 0
		first_v := 0
		second := 0
		second_v := 0

		for _, m := range musicList[g] {
			if plays[m] > first_v {
				second = first
				second_v = first_v
				first = m
				first_v = plays[m]
			} else if plays[m] > second_v {
				second = m
				second_v = plays[m]
			}
		}

		answer = append(answer, first)
		answer = append(answer, second)
	}

	return answer
}