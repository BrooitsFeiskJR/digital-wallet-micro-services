package validation

import "strconv"

func CheckCPF(strCPF string) bool {
	var Sum, Rest int
	Sum = 0

	if strCPF == "" {
		return false
	}

	if strCPF == "00000000000" {
		return false
	}

	if len(strCPF) < 11 || len(strCPF) > 11 {
		return false
	}

	for i := 1; i <= 9; i++ {
		digit, _ := strconv.Atoi(string(strCPF[i-1]))
		Sum += digit * (11 - i)
	}
	Rest = (Sum * 10) % 11

	if Rest == 10 || Rest == 11 {
		Rest = 0
	}
	if Rest != int(strCPF[9]-'0') {
		return false
	}

	Sum = 0
	for i := 1; i <= 10; i++ {
		digit, _ := strconv.Atoi(string(strCPF[i-1]))
		Sum += digit * (12 - i)
	}
	Rest = (Sum * 10) % 11

	if Rest == 10 || Rest == 11 {
		Rest = 0
	}
	if Rest != int(strCPF[10]-'0') {
		return false
	}

	return true
}
