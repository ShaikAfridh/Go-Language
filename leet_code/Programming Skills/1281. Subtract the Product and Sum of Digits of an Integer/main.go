func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	for n > 0 {
		product = product * (n % 10)
		sum = sum + (n % 10)
		n = n / 10
	}
	return product - sum
}