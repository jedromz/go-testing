package linting

func checkFlag(flag bool) bool {
	if flag {
		return true
	}
	return false
}
func errChecking() (int, error) {
	return 0, nil
}
func callsErrChecking() int {
	val, _ := errChecking()
	return val
}
