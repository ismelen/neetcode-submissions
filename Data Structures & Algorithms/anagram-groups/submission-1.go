import "slices"

func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[string][]string)

    for _, str := range strs {
        runes := []rune(str)
        slices.Sort(runes)
        key := string(runes)

        anagrams[key] = append(anagrams[key], str)
    }

    results := [][]string{}
    for _, v := range anagrams {
        results = append(results, v)
    }

    return results
}
