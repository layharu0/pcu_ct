package marathon

func marathonSolution(participant []string, completion []string) string {
	answer := ""
	check := make(map[string]bool, len(completion))

	for _, p := range participant {
		flag := false
		for _, c := range completion {
			val, exists := check[c]
			if !exists {
				check[c] = false
				val = check[c]
			}

			if !val && p == c {
				check[c] = true
				flag = true
				break
			}

		}

		if !flag {
			answer = p
			break
		}
	}

	return answer
}
