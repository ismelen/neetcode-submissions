type Solution struct{}

const SEP = "__"

func (s *Solution) Encode(strs []string) string {
    res := ""
    for _, str := range strs {
        res += fmt.Sprintf("%s%s", str, SEP)
    }
    return res
}

func (s *Solution) Decode(encoded string) []string {
    if encoded == "" {
        return []string{}
    }

    parts := strings.Split(encoded, SEP)
    return parts[:len(parts)-1]
}
