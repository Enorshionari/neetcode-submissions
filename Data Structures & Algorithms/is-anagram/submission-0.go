func isAnagram(s string, t string) bool {
    if len(s)!=len(t) { return false}
    ms := make(map[rune]int)
    mt := make(map[rune]int)
    for _,c := range s{
        ms[c]++
    }
    for _,c := range t{
        mt[c]++
    }
    for k,_ := range ms {
        if ms[k] != mt[k] {
            return false
        }
    }
    return true
}
