func isAnagram(s string, t string) bool {
	s_rune := []rune(s)
	t_rune := []rune(t)
	if len(s_rune) != len(t_rune) {
	    return false
	}
	count := make(map[rune]int)
	for i:=0; i<len(s_rune); i++ {
		count[s_rune[i]]++;
		count[t_rune[i]]--;
	}

	for _, value := range count {
		if value != 0 {
			return false
		}
	}
	return true
}
