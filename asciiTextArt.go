package asciitextart

// For now, This method only support a~z A~Z 0~9
// If you input a string with other characters(not include in a~z A~Z 0~9), this method will return ""
func ToAsciiTextArt(str string) string {
	res := lettersConver([]rune(str))
	return res
}

func lettersConver(b []rune) string {
	str := make([]string, 10)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			for j := 0; j < 10; j++ {
				str[j] += engLowerCase[int(b[i]-'a')*10+j]
			}
		} else if b[i] >= 'A' && b[i] <= 'Z' {
			for j := 0; j < 10; j++ {
				str[j] += engUpperCase[int(b[i]-'A')*10+j]
			}
		} else if b[i] >= '0' && b[i] <= '9' {
			for j := 0; j < 10; j++ {
				str[j] += numberCase[int(b[i]-'0')*10+j]
			}
		} else {
			// This branch should not be reached
			return ""
		}
	}
	var res string
	for i := 0; i < len(str); i++ {
		res += str[i] + "\r\n"
	}
	return res
}
