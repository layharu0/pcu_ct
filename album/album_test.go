package album

import "testing"

func TestSolution(t *testing.T) {
	genres := []string{"classic", "pop", "classic", "classic", "pop"}
	plays := []int{500, 600, 150, 800, 2500}
	result := []int{4, 1, 3, 0}

	for i := 0; i < 3; i++ {
		r := albumSolution(genres, plays)
		for i = 0; i < len(result); i++ {
			if r[i] != result[i] {
				t.Error("결과가 ", result, "가 아닙니다. 결과 : ", r)
			}
		}
	}
}
