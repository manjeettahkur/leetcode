package leetcode

//The two sum problem is a common interview question,
// and it is a variation of the subset sum problem.
// There is a popular dynamic programming solution for the subset sum problem,
// but for the two sum problem we can actually
// write an algorithm that runs in O(n) time.
// The challenge is to find all the pairs of two integers in an unsorted array
// that sum up to a given S.
func TwoSum(arr []int, s int) [][]int {
	var sum [][]int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == s {
				sum = append(sum, []int{arr[i], arr[j]})
			}
		}
	}
	return sum
}