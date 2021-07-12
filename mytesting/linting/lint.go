// https://golangci-lint.run/usage/install
package linting

func checkFlag(flag bool) bool {
	if flag == true {
		return true
	} else {
		return false
	}
}

func errChecking() (int, error) {
	return 0, nil
}

func callsErrChecking() int {
	val, _ := errChecking()
	return val
}
