package criteria

func countNumberofX(x int, numberCombination NumberCombination) int {
	count := 0
	if numberCombination.Blue == x {
		count++
	}
	if numberCombination.Yellow == x {
		count++
	}
	if numberCombination.Purple == x {
		count++
	}
	return count
}

func countEvenNumbers(numberCombination NumberCombination) int {
	count := 0
	if numberCombination.Blue%2 == 0 {
		count++
	}
	if numberCombination.Yellow%2 == 0 {
		count++
	}
	if numberCombination.Purple%2 == 0 {
		count++
	}
	return count
}

func countOddNumbers(numberCombination NumberCombination) int {
	count := 0
	if numberCombination.Blue%2 != 0 {
		count++
	}
	if numberCombination.Yellow%2 != 0 {
		count++
	}
	if numberCombination.Purple%2 != 0 {
		count++
	}
	return count
}

func numbersInAsecndingSequentialOrder(numberCombination NumberCombination) int {
	count := 0
	if numberCombination.Blue+1 == numberCombination.Yellow {
		count++
	}
	if numberCombination.Yellow+1 == numberCombination.Purple {
		count++
	}
	if count > 0 {
		count++
	}
	return count
}

func numbersInDescendingSequentialOrder(numberCombination NumberCombination) int {
	count := 0
	if numberCombination.Blue-1 == numberCombination.Yellow {
		count++
	}
	if numberCombination.Yellow-1 == numberCombination.Purple {
		count++
	}
	if count > 0 {
		count++
	}
	return count
}
