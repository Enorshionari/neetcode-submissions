func isAnagram(s string, t string) bool {
    if len(s)!=len(t) { return false}
    ms := make(map[rune]int)
    for _,c := range s{
        ms[c]++
    }
    for _,c := range t{
        ms[c]--
    }
    for k,_ := range ms {
        if ms[k] != 0 {
            return false
        }
    }
    return true
}
