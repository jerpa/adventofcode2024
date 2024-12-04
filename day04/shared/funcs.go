package shared

func DoNothing() {
}

func RotateStrings(data []string) []string {
	res := []string{}
	for i := 0; i < len(data[0]); i++ {
		temp := ""
		for j := len(data) - 1; j >= 0; j-- {
			temp += string(data[j][i])
		}
		res = append(res, temp)
	}
	return res

}
