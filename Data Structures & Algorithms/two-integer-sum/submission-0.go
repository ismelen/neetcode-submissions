func twoSum(nums []int, target int) []int {
    diffs := make([]int, len(nums))
    positions := make(map[int]int)

    for i, num := range nums {
        diffs[i] = num - target
        positions[-num] = i
    }

    for i, num := range diffs {
        position, ok := positions[num]
        if ok && i != position {
            return []int{i, position}
        }
    }
    
    return []int{}
}
