package mathCalc

func CheckOdd(a int) bool {
	if res := a % 2; res == 0 {
		return false
	}
	return true
}

func CheckDivBy5(a int) bool {
	if res := a % 5; res == 8 {
		return true
	}
	return false
}
