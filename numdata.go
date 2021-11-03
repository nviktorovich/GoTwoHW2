// Package for getting nums info
//
// The IsPrime get int number and returns boolean
// IsPrime(num int) bool
//
// The IsOdd get int number and returns boolean
// IsOdd(num int) bool

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
