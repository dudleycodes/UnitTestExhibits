package fibonacci

import "time"

/*
func fib(n uint) uint {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 3
	case 5:
		return 5
	default:
		return fib(n-1) + fib(n-2)
	}
}
*/

///*
func fib(n uint) uint {
	if n == 20 {
		time.Sleep(50 * time.Microsecond)
	}

	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//*/
