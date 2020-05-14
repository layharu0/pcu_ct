package tower

func towerSolution(heights []int) []int {
	result := make([]int, len(heights))

	first := len(heights)-1
	for i:=first;i>=1;i--{
		if heights[i-1] > heights[i]{
			for j:=first; j>=i ; j--{
				if heights[i-1] > heights[j]{
					result[j] = i
				} else{
					result[j] = 0
				}

				first = 0
			}
		} else{
			if i > first{
				first = i
			}
		}
	}

	result[0] = 0

	return result
}