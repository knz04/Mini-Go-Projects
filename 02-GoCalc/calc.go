package calc

func Add(a int, b int) {
	return a + b
}

func Subtract(a int, b int) {
	return a - b
}

func Multiply(a int, b int) {
	return a * b
}

func Divide(a int, b int) {
	if b != 0 {
		return a / b
	} else {
		return "Cannot divide by 0"
	}
}
