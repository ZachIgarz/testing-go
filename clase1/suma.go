package clase1

func Add(a, b int) int {
	return a + b
}

func AddMultiple(numbers ...int) int {
	valor := 0
	for _, v := range numbers {
		valor += v
	}
	return valor
}
