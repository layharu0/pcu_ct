package spy

import "testing"

func TestSolution(t *testing.T) {
	var clothes [][][]string
	clothes = [][][]string{
		{{"yello_hat", "headgear"}, {"blue_sunglasses", "eyewear"}, {"green_turban", "headgear"}},
		{{"crow_mask", "face"}, {"blue_sunglasses", "face"}, {"smoky_makeup", "face"}},
	}
	result := []int{5, 3}

	for i := 0; i < 2; i++ {
		r := spySolution(clothes[i])
		if r != result[i] {
			t.Error("결과가 ", result[i], "가 아닙니다. 결과 : ", r)
		}
	}
}
