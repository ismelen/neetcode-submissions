func hasDuplicate(nums []int) bool {
    visitedNums := make(map[int]bool)
    for _, num := range nums {
        _, ok := visitedNums[num]
        if !ok {
            visitedNums[num] = true
            continue
        }
        return true
    }
    return false
}
