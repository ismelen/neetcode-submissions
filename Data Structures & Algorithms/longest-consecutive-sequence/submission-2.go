func longestConsecutive(nums []int) int {
	maxSeq := 0
	prevs := make(map[int]bool)
	nexts := make(map[int]bool)

	for _, num := range nums {
		prevCount := 0
		for prevs[num-(prevCount+1)]{
			prevCount++
		}

		nextCount := 0
		for nexts[num+nextCount+1] {
			nextCount++
		}

		prevs[num] = true
		nexts[num] = true
		seq := 1 + prevCount + nextCount

		if seq > maxSeq {
			maxSeq = seq
		}
	}

	return maxSeq
}