package phone

import (
	"sort"
	"strings"
)

func phoneSolution(phone_book []string) bool {
	result := true
	sort.Strings(phone_book)

	for i := 0; i < len(phone_book)-1; i++ {
		for j := i + 1; j < len(phone_book); j++ {
			if strings.HasPrefix(phone_book[j], phone_book[i]) {
				result = false
				break
			}
		}

		if !result {
			break
		}
	}

	return result
}
