func topKFrequent(nums []int, k int) []int {
    counts := make(map[int]int)
    for _, num := range nums {
        counts[num]++
    }

    mostFrequent := []int{}
    for i:=0; i < k; i++ {
        max := -1
        res := -1
        for num, count := range counts {
            if count > max {
                max = count
                res = num
            }
        }
        delete(counts, res)
        mostFrequent = append(mostFrequent, res)
    }

    return mostFrequent
}
