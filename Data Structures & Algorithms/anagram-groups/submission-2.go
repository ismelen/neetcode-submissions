import "slices"

func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[string][]string)

    for _, str := range strs {
        runes := []rune(str)
        slices.Sort(runes)

        anagrams[string(runes)] = append(anagrams[string(runes)], str)
    }

    results := [][]string{}
    for _, v := range anagrams {
        results = append(results, v)
    }

    return results
}
