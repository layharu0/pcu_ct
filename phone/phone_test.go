package phone

import "testing"

func TestSolution(t *testing.T) {
	var phone_book [][]string
	phone_book = [][]string{{"119", "97674223", "1195524421"}, {"123", "456", "789"}, {"12", "123", "1235", "567", "88"}}
	result := []bool{false, true, false}

	for i := 0; i < 3; i++ {
		r := phoneSolution(phone_book[i])
		if r != result[i] {
			t.Error("결과가 ", result[i], "가 아닙니다. 결과 : ", r)
		}
	}
}
