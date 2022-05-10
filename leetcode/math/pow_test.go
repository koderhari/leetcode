package math

func myPow(x float64, n int) float64 {
	if n < 0 {
		n *= -1
		x = 1 / x
	}

	return doPow(x, n)
}

func doPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	y := doPow(x, n/2)

	if n%2 == 0 {
		return y * y
	} else {
		return x * y * y
	}
}
