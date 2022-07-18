package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tByte := []byte(t)
	sByte := []byte(s)
	sMap := map[byte]byte{}
	for i, b := range sByte {
		if _, ok := sMap[b]; !ok {
			sMap[b] = tByte[i]
		} else if sMap[b] != tByte[i] {
			return false
		}
	}
	return true
}
