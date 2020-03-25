package tool

func HashCode(s string) (hash int32) {
	if len(s) == 0 {
		return 0
	}
	hash = 0
	chr := 0
	for i := 0; i < len(s); i++ {
		chr = int(rune(s[i]))
		hash = ((hash << 5) - hash) + int32(chr)
		hash |= 0
	}
	return hash
}
