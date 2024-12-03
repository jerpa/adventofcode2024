package shared

import c "github.com/jerpa/adventofcode2024/common"

func DoNothing() {
}

func CheckLevel(s []string) bool {
	asc := false
	if c.GetInt(s[1]) > c.GetInt(s[0]) {
		asc = true
	}
	ok := true
	for i := 1; i < len(s); i++ {
		if asc && (c.GetInt(s[i])-c.GetInt(s[i-1]) < 1 || c.GetInt(s[i])-c.GetInt(s[i-1]) > 3) {
			ok = false
			break
		}
		if !asc && (c.GetInt(s[i])-c.GetInt(s[i-1]) > -1 || c.GetInt(s[i])-c.GetInt(s[i-1]) < (-3)) {
			ok = false
			break
		}
	}
	return ok
}
