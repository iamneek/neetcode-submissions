func hasDuplicate(nums []int) bool {
    mp := make(map[int]int)
    for _, n := range nums {
        if _, ok := mp[n]; ok {
            return true
        } else {
            mp[n] = n
        }
    }
    return false
}
