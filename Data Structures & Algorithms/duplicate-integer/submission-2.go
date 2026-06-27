func hasDuplicate(nums []int) bool {
    if len(nums) < 2 {
        return false
    }
   another := make([]int, 0,len(nums))
   for _, num:= range nums {
    for _, n := range another {
        if num == n {
            return true
        }
    }
    another = append(another, num)
   }
   return false
}
