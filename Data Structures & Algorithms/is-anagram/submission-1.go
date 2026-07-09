import (
    "slices"
)

func isAnagram(s string, t string) bool {
    sortedS := []rune(s)
    sortedT := []rune(t)
    
    slices.Sort(sortedS)
    slices.Sort(sortedT)

    return string(sortedS) == string(sortedT)
}
