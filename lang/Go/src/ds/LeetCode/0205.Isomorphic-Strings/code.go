package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := map[rune]byte{}
	for i, b := range s {
		if _, ok := sMap[b]; !ok {
			sMap[b] = t[i]
		} else if sMap[b] != t[i] {
			return false
		}
	}
	return true
}
