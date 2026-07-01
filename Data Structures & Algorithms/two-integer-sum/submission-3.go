func twoSum(nums []int, target int) []int {
	seenmp := make(map[int]int)
	for idx, num := range nums {
		diff := target - num
		if v, ok := seenmp[diff]; ok {
			return []int{v, idx}
		}
		seenmp[num] = idx
	}
	return []int{}
}
