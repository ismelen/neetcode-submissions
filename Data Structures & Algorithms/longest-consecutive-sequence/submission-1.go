func longestConsecutive(nums []int) int {
	maxSeq := 0
	prevs := make(map[int]int)
	nexts := make(map[int]int)

	for _, num := range nums {
		prevCount := 0
		for {
			seq := prevs[num-(prevCount+1)]
			if seq == 0 {
				break
			}
			prevCount++
		}

		nextCount := 0
		for {
			seq := nexts[num+nextCount+1]
			if seq == 0 {
				break
			}
			nextCount++
		}

		prevs[num] = 1
		nexts[num] = 1
		seq := 1 + prevCount + nextCount

		if seq > maxSeq {
			maxSeq = seq
		}
	}

	return maxSeq
}