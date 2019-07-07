package leetcode

func TwoSum1(nums []int, target int) [][]int {
	var r = make([][]int, 0)
	var m = make(map[int]int, 0)
	for p, i := range nums {
		j := target - i
		if pos, ok := m[j]; ok {
			if pos != p {
				r = append(r, []int{j, i})
			}
		} else {
			m[i] = p
		}
	}
	return r
}

