package marathon

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var participant [][]string
	var completion [][]string

	participant = [][]string{{"leo", "kiki", "eden"}, {"marina", "josipa", "nikola", "vinko", "fillipa"}, {"mislav", "stanko", "mislav", "ana"}}
	completion = [][]string{{"eden", "kiki"}, {"josipa", "fillipa", "marina", "nikola"}, {"stanko", "ana", "mislav"}}

	result := [...]string{"leo", "vinko", "mislav"}
	for i := 0; i < 3; i++ {
		r := marathonSolution(participant[i], completion[i])
		if r != result[i] {
			t.Error("결과가 ", result[i], "가 아닙니다. 결과 : ", r)
		}
	}

}
