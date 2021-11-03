package GoTwoHW2

// IsPrime function return true if number is prime.
func IsPrime(num int) bool {
	var flag bool
	if num == 1 {
		return true
	}
	for i := 2; i < num/2+1; i++ {
		if num%i == 0 {
			return flag
		}
	}
	return !flag
}

// IsOdd function return true if number is odd.
func IsOdd(num int) bool {
	return num%2 != 0
}
