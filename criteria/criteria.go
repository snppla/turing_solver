package criteria

import "fmt"

type Verifier struct {
	TestCombination func(numberCombination NumberCombination) bool
}

// NumberCombination is a struct that represents the combination of numbers between 1 and 5
type NumberCombination struct {
	Blue   int
	Yellow int
	Purple int
}

func (nc NumberCombination) ToString() string {
	return fmt.Sprintf("Blue: %d, Yellow: %d, Purple: %d", nc.Blue, nc.Yellow, nc.Purple)
}

type Criteria struct {
	GetVerifiers func() []Verifier
}

type BaseCriteria struct {
	verifiers []Verifier
}

func (c BaseCriteria) GetVerifiers() []Verifier {
	return c.verifiers
}

func GetCriteriaList() []Criteria {
	return append([]Criteria{}, criteriaList...)
}

var criteriaList = []Criteria{
	// there is no criteria 0
	{},
	{
		// criteria1 checks if the blue number is 1
		// or if the blue number is greater than 1
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > 1
					},
				},
			}
		},
	},
	{
		// criteria2 checks if blue is less than 3
		// or blue is equal to 3
		// or blue is greater than 3
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > 3
					},
				},
			}
		},
	},
	{
		// criteria 3 checks if the yellow number is less than 3
		// or yellow number is equal to 3
		// or yellow number is greater than 3
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > 3
					},
				},
			}
		},
	},
	{
		// criteria 4 checks if yellow is less than 4, equal to 4, or greater than 4
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > 4
					},
				},
			}
		},
	},
	{
		// criteria 5 checks if blue is even or odd
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue%2 == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue%2 != 0
					},
				},
			}
		},
	},
	{
		// criteria 6 checks if yellow is even or odd
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow%2 == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow%2 != 0
					},
				},
			}
		},
	},
	{

		// criteria 7 checks if purple is even or odd
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple%2 == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple%2 != 0
					},
				},
			}
		},
	},
	{
		// criteria 8  checks if there are no 1s, a single 1, two 1s, or three 1s
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(1, numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(1, numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(1, numberCombination) == 2
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(1, numberCombination) == 3
					},
				},
			}
		},
	},
	{
		// criteria 9 checks if there are no 3s, a single 3, two 3s, or three 3s
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(3, numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(3, numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(3, numberCombination) == 2
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(3, numberCombination) == 3

					},
				},
			}
		},
	},
	{
		// criteria 10 checks if there are no 4s, a single 4, two 4s, or three 4s
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(4, numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(4, numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(4, numberCombination) == 2
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(4, numberCombination) == 3
					},
				},
			}
		},
	},
	{
		// criteria 11 checks if blue is less than yellow, blue is equal to yellow, or blue is greater than yellow
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < numberCombination.Yellow
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == numberCombination.Yellow
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > numberCombination.Yellow
					},
				},
			}
		},
	},
	{
		// criteria 12 checks if blue is less than purple, blue is equal to purple, or blue is greater than purple
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > numberCombination.Purple
					},
				},
			}
		},
	},
	{
		// criteria 13 checks if yellow is less than purple, yellow is equal to purple, or yellow is greater than purple
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > numberCombination.Purple
					},
				},
			}
		},
	},
	{
		// criteria 14 checks if blue is the smallest number, yellow is the smallest number, or purple is the smallest number
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < numberCombination.Yellow && numberCombination.Blue < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < numberCombination.Blue && numberCombination.Yellow < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple < numberCombination.Blue && numberCombination.Purple < numberCombination.Yellow
					},
				},
			}
		},
	},
	{
		// criteria 15 checks if blue is the largest number, yellow is the largest number, or purple is the largest number
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > numberCombination.Yellow && numberCombination.Blue > numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > numberCombination.Blue && numberCombination.Yellow > numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple > numberCombination.Blue && numberCombination.Purple > numberCombination.Yellow
					},
				},
			}
		},
	},
	{
		// criteria 16 checks if there are more even numbers than odd numbers, or more odd numbers than even numbers
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countEvenNumbers(numberCombination) > countOddNumbers(numberCombination)
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countEvenNumbers(numberCombination) < countOddNumbers(numberCombination)
					},
				},
			}
		},
	},
	{
		// criteria 17 checks if there are 0 even numbers, 1 even number, 2 even numbers, or 3 even numbers
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countEvenNumbers(numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countEvenNumbers(numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countEvenNumbers(numberCombination) == 2
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countEvenNumbers(numberCombination) == 3
					},
				},
			}
		},
	},
	{
		// criteria 18 checks if the sum of the numbers is even or odd
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return (numberCombination.Blue+numberCombination.Yellow+numberCombination.Purple)%2 == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return (numberCombination.Blue+numberCombination.Yellow+numberCombination.Purple)%2 != 0
					},
				},
			}
		},
	},
	{
		// criteria 19 checks if the sum of blue and yellow is less than 6, equal to 6, or greater than 6
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue+numberCombination.Yellow < 6
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue+numberCombination.Yellow == 6
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue+numberCombination.Yellow > 6
					},
				},
			}
		},
	},
	{
		// criteria 20 checks if a number repeats 3 times, 2 times, or 0 times
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(numberCombination.Blue, numberCombination) == 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(numberCombination.Blue, numberCombination) == 2 ||
							countNumberofX(numberCombination.Yellow, numberCombination) == 2 ||
							countNumberofX(numberCombination.Purple, numberCombination) == 2
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(numberCombination.Blue, numberCombination) == 1 &&
							countNumberofX(numberCombination.Yellow, numberCombination) == 1 &&
							countNumberofX(numberCombination.Purple, numberCombination) == 1
					},
				},
			}
		},
	},
	{
		// criteria 21 checks that there is no pairs or 1 pair
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(numberCombination.Blue, numberCombination) == 1 &&
							countNumberofX(numberCombination.Yellow, numberCombination) == 1 &&
							countNumberofX(numberCombination.Purple, numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(numberCombination.Blue, numberCombination) == 2 ||
							countNumberofX(numberCombination.Yellow, numberCombination) == 2 ||
							countNumberofX(numberCombination.Purple, numberCombination) == 2
					},
				},
			}
		},
	},
	{
		// criteria 22 checks the 3 numbers are in ascending order, descending order, or any other order
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < numberCombination.Yellow && numberCombination.Yellow < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > numberCombination.Yellow && numberCombination.Yellow > numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return !(numberCombination.Blue < numberCombination.Yellow && numberCombination.Yellow < numberCombination.Purple) &&
							!(numberCombination.Blue > numberCombination.Yellow && numberCombination.Yellow > numberCombination.Purple)
					},
				},
			}
		},
	},
	{
		// criteria 23 checks if the sum of all numbers is less than 6, equal to 6, or greater than 6
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue+numberCombination.Yellow+numberCombination.Purple < 6
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue+numberCombination.Yellow+numberCombination.Purple == 6
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue+numberCombination.Yellow+numberCombination.Purple > 6
					},
				},
			}
		},
	},
	{
		// criteria 24 checks if the 3 numbers are in sequential ascending order, 2 are in sequential ascending order, or neither
		GetVerifiers: func() []Verifier {

			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numbersInAsecndingSequentialOrder(numberCombination) == 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numbersInAsecndingSequentialOrder(numberCombination) == 2
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numbersInAsecndingSequentialOrder(numberCombination) == 0
					},
				},
			}
		},
	},
	{
		// criteria 25 checks if there is no sequence of asecnding or descending numbers,
		// if there are 2 numbers in a sequence of ascending or descending numbers,
		// or if there are 3 numbers in a sequence of ascending or descending numbers
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numbersInAsecndingSequentialOrder(numberCombination) == 0 && numbersInDescendingSequentialOrder(numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numbersInAsecndingSequentialOrder(numberCombination) == 2 || numbersInDescendingSequentialOrder(numberCombination) == 2
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numbersInAsecndingSequentialOrder(numberCombination) == 3 || numbersInDescendingSequentialOrder(numberCombination) == 3
					},
				},
			}
		},
	},
	{
		// criteria 26 checks that a specific color is less than 3
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple < 3
					},
				},
			}
		},
	},
	{
		// criteria 27 checks that a specific color is less than 4
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple < 4
					},
				},
			}
		},
	},
	{
		// criteria 28 checks that a specific color is equal to 1
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple == 1
					},
				},
			}
		},
	},
	{
		// criteria 29 checks that a specific color is equal to 3
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple == 3
					},
				},
			}
		},
	},
	{
		// critera 30 checks that a specific color is equal to 4
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple == 4
					},
				},
			}
		},
	},
	{
		// criteria 31 checks that a specific color is greater than 1
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple > 1
					},
				},
			}
		},
	},
	{
		// criteria 32 checks that a specific color is greater than 3
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple > 3
					},
				},
			}
		},
	},
	{
		// criteria 33 checks that a specific color is even or odd
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue%2 == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow%2 == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple%2 == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue%2 != 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow%2 != 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple%2 != 0
					},
				},
			}
		},
	},
	{
		// criteria 34 checks that a specific color is the smallest
		// it can also be equal with another color
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue <= numberCombination.Yellow && numberCombination.Blue <= numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow <= numberCombination.Blue && numberCombination.Yellow <= numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple <= numberCombination.Blue && numberCombination.Purple <= numberCombination.Yellow
					},
				},
			}
		},
	},
	{
		// criteria 35 checks that a specific color is the largest
		// it can also be equal with another color
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue >= numberCombination.Yellow && numberCombination.Blue >= numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow >= numberCombination.Blue && numberCombination.Yellow >= numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple >= numberCombination.Blue && numberCombination.Purple >= numberCombination.Yellow
					},
				},
			}
		},
	},
	{
		// criteria 36 checks that the sum of all 3 colors is a multiple of 3, 4, or 5
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return (numberCombination.Blue+numberCombination.Yellow+numberCombination.Purple)%3 == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return (numberCombination.Blue+numberCombination.Yellow+numberCombination.Purple)%4 == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return (numberCombination.Blue+numberCombination.Yellow+numberCombination.Purple)%5 == 0
					},
				},
			}
		},
	},
	{
		// criteria 37 checks that the sum of two colors equals 4
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue+numberCombination.Yellow == 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue+numberCombination.Purple == 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow+numberCombination.Purple == 4
					},
				},
			}
		},
	},
	{
		// criteria 38 checks that the sum of two colors equals 6
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue+numberCombination.Yellow == 6
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue+numberCombination.Purple == 6
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow+numberCombination.Purple == 6
					},
				},
			}
		},
	},
	{
		// criteria 39 compares one specific color to 1
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple > 1
					},
				},
			}
		},
	},
	{
		//criteria 40 compares a specific color to 3
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple < 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple == 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > 3
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple > 3
					},
				},
			}
		},
	},
	{
		// criteria 41 compares a specific color to 4
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple < 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple == 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > 4
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple > 4
					},
				},
			}
		},
	},
	{
		// critiera 42 checks that a specific color is the smallest or largest
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < numberCombination.Yellow && numberCombination.Blue < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < numberCombination.Blue && numberCombination.Yellow < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple < numberCombination.Blue && numberCombination.Purple < numberCombination.Yellow
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > numberCombination.Yellow && numberCombination.Blue > numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > numberCombination.Blue && numberCombination.Yellow > numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Purple > numberCombination.Blue && numberCombination.Purple > numberCombination.Yellow
					},
				},
			}
		},
	},
	{
		// criteria 43 compares blue to another color
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < numberCombination.Yellow
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == numberCombination.Yellow
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > numberCombination.Yellow
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > numberCombination.Purple
					},
				},
			}
		},
	},
	{
		// criteria 44 compares yellow to another color
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < numberCombination.Blue
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == numberCombination.Blue
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > numberCombination.Blue
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > numberCombination.Purple
					},
				},
			}
		},
	},
	{
		// criteria 45 tests how many 1's or 3's are in the combination
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(1, numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(1, numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(1, numberCombination) == 2
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(3, numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(3, numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(3, numberCombination) == 2
					},
				},
			}
		},
	},
	{
		// criteria 46 checks how many 3's or 4's are in the combination
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(3, numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(3, numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(3, numberCombination) == 2
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(4, numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(4, numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(4, numberCombination) == 2
					},
				},
			}
		},
	},
	{
		// criteria 47 checks how many 1's or 4's are in the combination
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(1, numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(1, numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(1, numberCombination) == 2
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(4, numberCombination) == 0
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(4, numberCombination) == 1
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return countNumberofX(4, numberCombination) == 2
					},
				},
			}
		},
	},
	{
		// criteria 48 compares one specific color to another specific color
		GetVerifiers: func() []Verifier {
			return []Verifier{
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < numberCombination.Yellow
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow < numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == numberCombination.Yellow
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue == numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow == numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > numberCombination.Yellow
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Blue > numberCombination.Purple
					},
				},
				{
					TestCombination: func(numberCombination NumberCombination) bool {
						return numberCombination.Yellow > numberCombination.Purple
					},
				},
			}
		},
	},
}
