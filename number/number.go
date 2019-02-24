package number

// GCD : Return the greatest common divisor
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
