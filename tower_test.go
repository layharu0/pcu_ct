package tower

import "testing"

func TestSolution(t *testing.T) {
	var heights [][]int
	var result [][]int

	heights = [][]int{{6, 9, 5, 7, 4}, {3, 9, 9, 3, 5, 7, 2}, {1, 5, 3, 6, 7, 6, 5}}
	result = [][]int{{0, 0, 2, 2, 4}, {0, 0, 0, 3, 3, 3, 6}, {0, 0, 2, 0, 0, 5, 6}}

	for i := 0; i < 2; i++ {
		r := towerSolution(heights[i])
		for j := 0; j < len(result[i]); j++ {
			if r[j] != result[i][j] {
				t.Error("정답이 일치하지 않습니다. 정답 : ", result[i], ", 답안 : ", r)
			}
		}
	}
}
