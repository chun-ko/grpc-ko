package race

func Size(text string) int {
	return len(text)
}

func Size2(text string) int {
	return len(text)
}

func Size3(text string) int {
	count := 0
	for i := 0; i < 10; i++ {
		count++
	}
	for i := 0; i < 10; i++ {
		count++
	}

	for i := 0; i < 10; i++ {
		count++
	}
	for i := 0; i < 10; i++ {
		count++
	}
	for i := 0; i < 10; i++ {
		count++
	}
	for i := 0; i < 10; i++ {
		count++
	}
	for i := 0; i < 10; i++ {
		count++
	}

	return len(text)
}
