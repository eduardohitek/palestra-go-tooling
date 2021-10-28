package main

func isOdd(number int32) bool {
	rest := number % 2
	if rest == 0 {
		return true
	}
	return false
}
