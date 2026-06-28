import "maps"

func isAnagram(s string, t string) bool {
	s_mp := make(map[string]int)
	t_mp := make(map[string]int)
	if len(s) != len(t) {
	    return false
	}
	for i, b:= range s {
        if _, ok:= s_mp[string(b)]; ok {
            s_mp[string(b)]++
        } else {
	        s_mp[string(b)] = 1
        }
        if _, ok:= t_mp[string(t[i])]; ok {
            t_mp[string(t[i])]++
        } else {
	        t_mp[string(t[i])] = 1
        }
	}
	return maps.Equal(s_mp, t_mp)
}
