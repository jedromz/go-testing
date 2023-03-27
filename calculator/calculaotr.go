package calculator

// CalculateIsArmstrong is a function to check if a number is armstrong or not
func CalculateIsArmstrong(n int) bool {
	var sum int
	for i := n; i > 0; i /= 10 {
		sum += (i % 10) * (i % 10) * (i % 10)
	}
	return sum == n
}
func GenerateArmstrongNumbers() []int {
	var armstrongNumbers []int
	for i := 100; i < 1000; i++ {
		if CalculateIsArmstrong(i) {
			armstrongNumbers = append(armstrongNumbers, i)
		}
	}
	return armstrongNumbers
}
