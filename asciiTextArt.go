package asciitextart

// This method only support a~z A-Z 0~9
func ToAsciiTextArt(str string) string {
	res := lettersConver([]byte(str))
	return res
}

func lettersConver(b []byte) string {
	str := make([]string, 10)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			for j := 0; j < 10; j++ {
				str[j] += lowercase[int(b[i]-'a')*10+j]
			}
		}
	}
	var res string
	for i := 0; i < len(str); i++ {
		res += str[i] + "\r\n"
	}
	return res
}
