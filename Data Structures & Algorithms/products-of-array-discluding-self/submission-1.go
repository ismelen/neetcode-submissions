func productExceptSelf(nums []int) []int {
    total := 1
    zeroCount := 0
    for _, num := range nums {
        if num == 0 {
            zeroCount++
            continue
        }
        total *= num 
    }

    if zeroCount > 1 {
        total = 0
    }

    res := make([]int, len(nums))
    if total == 0 {
        return res
    }

    for i:=0; i < len(nums); i++ {
        if nums[i] == 0 {
            res[i] = total
            continue
        }
        if zeroCount > 0 {
            res[i] = 0
            continue
        }
        res[i] = total / nums[i]
    }
    return res
}
