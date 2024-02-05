package criteria

import (
	"testing"
)

// Criteria 1 checks if the blue number is 1 or if the blue number is greater than 1
func TestCriteria1(t *testing.T) {
	criteria := criteriaList[1]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 2 {
		t.Errorf("Expected 2 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if blue == 1 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria1 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue != 1 && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria1 verifier 0 failed with combinationf %v", NumberCombination{blue, yellow, purple})
				}
				if blue > 1 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria1 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue <= 1 && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria1 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// Criteria 2 checks if blue is less than 3, equal to 3, or greater than 3
func TestCriteria2(t *testing.T) {
	criteria := criteriaList[2]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if blue < 3 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria2 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue >= 3 && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria2 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if blue == 3 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria2 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue != 3 && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria2 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if blue > 3 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria2 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue <= 3 && verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria2 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// Criteria 3 checks if yellow is less than 3, equal to 3, or greater than 3
func TestCriteria3(t *testing.T) {
	criteria := criteriaList[3]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if yellow < 3 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria3 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow >= 3 && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria3 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if yellow == 3 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria3 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow != 3 && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria3 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if yellow > 3 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria3 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow <= 3 && verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria3 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 4 checks if yellow is less than 4, equal to 4, or greater than 4
func TestCriteria4(t *testing.T) {
	criteria := criteriaList[4]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if yellow < 4 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria4 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow >= 4 && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria4 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if yellow == 4 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria4 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow != 4 && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria4 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if yellow > 4 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria4 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow <= 4 && verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria4 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 5 checks if blue is even or odd
func TestCriteria5(t *testing.T) {
	criteria := criteriaList[5]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 2 {
		t.Errorf("Expected 2 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if blue%2 == 0 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria5 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue%2 != 0 && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria5 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if blue%2 != 0 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria5 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue%2 == 0 && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria5 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 6 checks if yellow is even or odd
func TestCriteria6(t *testing.T) {
	criteria := criteriaList[6]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 2 {
		t.Errorf("Expected 2 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if yellow%2 == 0 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria6 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow%2 != 0 && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria6 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if yellow%2 != 0 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria6 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow%2 == 0 && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria6 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 7 checks if purple is even or odd
func TestCriteria7(t *testing.T) {
	criteria := criteriaList[7]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 2 {
		t.Errorf("Expected 2 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if purple%2 == 0 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria7 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if purple%2 != 0 && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria7 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if purple%2 != 0 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria7 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if purple%2 == 0 && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria7 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 8  checks if there are no 1s, a single 1, two 1s, or three 1s
func TestCriteria8(t *testing.T) {
	criteria := criteriaList[8]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 4 {
		t.Errorf("Expected 4 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				count := 0
				if combination.Blue == 1 {
					count++
				}
				if combination.Yellow == 1 {
					count++
				}
				if combination.Purple == 1 {
					count++
				}
				if count == 0 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 0 failed with combination %v", combination)
				} else if count != 0 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 0 failed with combination %v", combination)
				}

				if count == 1 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 1 failed with combination %v", combination)
				} else if count != 1 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 1 failed with combination %v", combination)
				}

				if count == 2 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 2 failed with combination %v", combination)
				} else if count != 2 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 2 failed with combination %v", combination)
				}

				if count == 3 && !verifiers[3].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 3 failed with combination %v", combination)
				} else if count != 3 && verifiers[3].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 3 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 9 checks if there are no 3s, a single 3, two 3s, or three 3s
func TestCriteria9(t *testing.T) {
	criteria := criteriaList[9]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 4 {
		t.Errorf("Expected 4 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				count := 0
				if combination.Blue == 3 {
					count++
				}
				if combination.Yellow == 3 {
					count++
				}
				if combination.Purple == 3 {
					count++
				}
				if count == 0 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria9 verifier 0 failed with combination %v", combination)
				} else if count != 0 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria9 verifier 0 failed with combination %v", combination)
				}

				if count == 1 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria9 verifier 1 failed with combination %v", combination)
				} else if count != 1 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria9 verifier 1 failed with combination %v", combination)
				}

				if count == 2 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria9 verifier 2 failed with combination %v", combination)
				} else if count != 2 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria9 verifier 2 failed with combination %v", combination)
				}

				if count == 3 && !verifiers[3].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 3 failed with combination %v", combination)
				} else if count != 3 && verifiers[3].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 3 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 10 checks if there are no 4s, a single 4, two 4s, or three 4s
func TestCriteria10(t *testing.T) {
	criteria := criteriaList[10]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 4 {
		t.Errorf("Expected 4 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				count := 0
				if combination.Blue == 4 {
					count++
				}
				if combination.Yellow == 4 {
					count++
				}
				if combination.Purple == 4 {
					count++
				}
				if count == 0 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria10 verifier 0 failed with combination %v", combination)
				} else if count != 0 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria10 verifier 0 failed with combination %v", combination)
				}

				if count == 1 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria10 verifier 1 failed with combination %v", combination)
				} else if count != 1 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria10 verifier 1 failed with combination %v", combination)
				}

				if count == 2 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria10 verifier 2 failed with combination %v", combination)
				} else if count != 2 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria10 verifier 2 failed with combination %v", combination)
				}

				if count == 3 && !verifiers[3].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 3 failed with combination %v", combination)
				} else if count != 3 && verifiers[3].TestCombination(combination) {
					t.Errorf("Criteria8 verifier 3 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 11 checks if blue is less than yellow, blue is equal to yellow, or blue is greater than yellow
func TestCriteria11(t *testing.T) {
	criteria := criteriaList[11]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if blue < yellow && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria11 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue >= yellow && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria11 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if blue == yellow && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria11 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue != yellow && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria11 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if blue > yellow && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria11 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue <= yellow && verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria11 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 12 checks if blue is less than purple, blue is equal to purple, or blue is greater than purple
func TestCriteria12(t *testing.T) {
	criteria := criteriaList[12]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if blue < purple && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria12 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue >= purple && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria12 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if blue == purple && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria12 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue != purple && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria12 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if blue > purple && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria12 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if blue <= purple && verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria12 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 13 checks if yellow is less than purple, yellow is equal to purple, or yellow is greater than purple
func TestCriteria13(t *testing.T) {
	criteria := criteriaList[13]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if yellow < purple && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria13 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow >= purple && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria13 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if yellow == purple && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria13 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow != purple && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria13 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if yellow > purple && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria13 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				} else if yellow <= purple && verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria13 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 14 checks if blue is the smallest number, yellow is the smallest number, or purple is the smallest number
func TestCriteria14(t *testing.T) {
	criteria := criteriaList[14]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				blueSmallest := blue < yellow && blue < purple
				yellowSmallest := yellow < blue && yellow < purple
				purpleSmallest := purple < blue && purple < yellow

				// make sure only 1 color is the smallest
				if blueSmallest && yellowSmallest {
					t.Errorf("Both blue and yellow are the smallest")
				}
				if blueSmallest && purpleSmallest {
					t.Errorf("Both blue and purple are the smallest")
				}
				if yellowSmallest && purpleSmallest {
					t.Errorf("Both yellow and purple are the smallest")
				}

				if blueSmallest && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria14 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if !blueSmallest && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria14 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if yellowSmallest && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria14 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if !yellowSmallest && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria14 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if purpleSmallest && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria14 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if !purpleSmallest && verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria14 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

			}
		}
	}
}

// criteria 15 checks if blue is the largest number, yellow is the largest number, or purple is the largest number
func TestCriteria15(t *testing.T) {
	criteria := criteriaList[15]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				blueLargest := blue > yellow && blue > purple
				yellowLargest := yellow > blue && yellow > purple
				purpleLargest := purple > blue && purple > yellow

				// make sure only 1 color is the largest
				if blueLargest && yellowLargest {
					t.Errorf("Both blue and yellow are the largest")
				}
				if blueLargest && purpleLargest {
					t.Errorf("Both blue and purple are the largest")
				}
				if yellowLargest && purpleLargest {
					t.Errorf("Both yellow and purple are the largest")
				}

				if blueLargest && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria15 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if !blueLargest && verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria15 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if yellowLargest && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria15 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if !yellowLargest && verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria15 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if purpleLargest && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria15 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if !purpleLargest && verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria15 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 16 checks if there are more even numbers than odd numbers, or more odd numbers than even numbers
func TestCriteria16(t *testing.T) {
	criteria := criteriaList[16]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 2 {
		t.Errorf("Expected 2 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				evenCount := 0
				oddCount := 0
				if combination.Blue%2 == 0 {
					evenCount++
				} else {
					oddCount++
				}
				if combination.Yellow%2 == 0 {
					evenCount++
				} else {
					oddCount++
				}
				if combination.Purple%2 == 0 {
					evenCount++
				} else {
					oddCount++
				}

				if evenCount > oddCount && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria16 verifier 0 failed with combination %v", combination)
				}
				if evenCount <= oddCount && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria16 verifier 0 failed with combination %v", combination)
				}

				if oddCount > evenCount && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria16 verifier 1 failed with combination %v", combination)
				}
				if oddCount <= evenCount && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria16 verifier 1 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 17 checks if there are 0 even numbers, 1 even number, 2 even numbers, or 3 even numbers
func TestCriteria17(t *testing.T) {
	criteria := criteriaList[17]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 4 {
		t.Errorf("Expected 4 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				evenCount := 0
				if combination.Blue%2 == 0 {
					evenCount++
				}
				if combination.Yellow%2 == 0 {
					evenCount++
				}
				if combination.Purple%2 == 0 {
					evenCount++
				}

				if evenCount == 0 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria17 verifier 0 failed with combination %v", combination)
				}
				if evenCount != 0 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria17 verifier 0 failed with combination %v", combination)
				}

				if evenCount == 1 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria17 verifier 1 failed with combination %v", combination)
				}
				if evenCount != 1 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria17 verifier 1 failed with combination %v", combination)
				}

				if evenCount == 2 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria17 verifier 2 failed with combination %v", combination)
				}
				if evenCount != 2 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria17 verifier 2 failed with combination %v", combination)
				}

				if evenCount == 3 && !verifiers[3].TestCombination(combination) {
					t.Errorf("Criteria17 verifier 3 failed with combination %v", combination)
				}
				if evenCount != 3 && verifiers[3].TestCombination(combination) {
					t.Errorf("Criteria17 verifier 3 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 18 checks if the sum of the numbers is even or odd
func TestCriteria18(t *testing.T) {
	criteria := criteriaList[18]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 2 {
		t.Errorf("Expected 2 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				sum := combination.Blue + combination.Yellow + combination.Purple
				if sum%2 == 0 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria18 verifier 0 failed with combination %v", combination)
				}
				if sum%2 != 0 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria18 verifier 0 failed with combination %v", combination)
				}

				if sum%2 != 0 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria18 verifier 1 failed with combination %v", combination)
				}
				if sum%2 == 0 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria18 verifier 1 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 19 checks if the sumb of blue and yellow is less than 6, equal to 6, or greater than 6
func TestCriteria19(t *testing.T) {
	criteria := criteriaList[19]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				sum := combination.Blue + combination.Yellow
				if sum < 6 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria19 verifier 0 failed with combination %v", combination)
				}
				if sum >= 6 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria19 verifier 0 failed with combination %v", combination)
				}

				if sum == 6 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria19 verifier 1 failed with combination %v", combination)
				}
				if sum != 6 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria19 verifier 1 failed with combination %v", combination)
				}

				if sum > 6 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria19 verifier 2 failed with combination %v", combination)
				}
				if sum <= 6 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria19 verifier 2 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 20 checks if a number repeats 3 times, 2 times, or 0 times
func TestCriteria20(t *testing.T) {
	criteria := criteriaList[20]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				count := 0
				if combination.Blue == combination.Yellow && combination.Blue == combination.Purple {
					count = 3
				} else if combination.Blue == combination.Yellow || combination.Blue == combination.Purple || combination.Yellow == combination.Purple {
					count = 2
				} else {
					count = 0
				}

				if count == 3 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria20 verifier 0 failed with combination %v", combination)
				}
				if count != 3 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria20 verifier 0 failed with combination %v", combination)
				}

				if count == 2 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria20 verifier 1 failed with combination %v", combination)
				}
				if count != 2 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria20 verifier 1 failed with combination %v", combination)
				}

				if count == 0 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria20 verifier 2 failed with combination %v", combination)
				}
				if count != 0 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria20 verifier 2 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 21 checks that there is no pairs or 1 pair
func TestCriteria21(t *testing.T) {
	criteria := criteriaList[21]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 2 {
		t.Errorf("Expected 2 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				count := 0
				if combination.Blue == combination.Yellow {
					count++
				}
				if combination.Blue == combination.Purple {
					count++
				}
				if combination.Yellow == combination.Purple {
					count++
				}

				if count == 0 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria21 verifier 0 failed with combination %v", combination)
				}
				if count != 0 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria21 verifier 0 failed with combination %v", combination)
				}

				if count == 1 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria21 verifier 1 failed with combination %v", combination)
				}
				if count != 1 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria21 verifier 1 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 22 checks the 3 numbers are in ascending order, descending order, or any other order
func TestCriteria22(t *testing.T) {
	criteria := criteriaList[22]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				if combination.Blue < combination.Yellow && combination.Yellow < combination.Purple {
					if !verifiers[0].TestCombination(combination) {
						t.Errorf("Criteria22 verifier 0 failed with combination %v", combination)
					}
				} else if combination.Blue > combination.Yellow && combination.Yellow > combination.Purple {
					if !verifiers[1].TestCombination(combination) {
						t.Errorf("Criteria22 verifier 1 failed with combination %v", combination)
					}
				} else {
					if !verifiers[2].TestCombination(combination) {
						t.Errorf("Criteria22 verifier 2 failed with combination %v", combination)
					}
				}
			}
		}
	}
}

// criteria 23 checks if the sum of all numbers is less than 6, equal to 6, or greater than 6
func TestCriteria23(t *testing.T) {
	criteria := criteriaList[23]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				sum := combination.Blue + combination.Yellow + combination.Purple
				if sum < 6 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria23 verifier 0 failed with combination %v", combination)
				}
				if sum >= 6 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria23 verifier 0 failed with combination %v", combination)
				}

				if sum == 6 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria23 verifier 1 failed with combination %v", combination)
				}
				if sum != 6 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria23 verifier 1 failed with combination %v", combination)
				}

				if sum > 6 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria23 verifier 2 failed with combination %v", combination)
				}
				if sum <= 6 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria23 verifier 2 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 24 checks if the 3 numbers are in sequential ascending order, 2 are in sequential ascending order, or neither
func TestCriteria24(t *testing.T) {
	criteria := criteriaList[24]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				if combination.Blue+1 == combination.Yellow && combination.Yellow+1 == combination.Purple {
					if !verifiers[0].TestCombination(combination) {
						t.Errorf("Criteria24 verifier 0 failed with combination %v", combination)
					}
				} else if combination.Blue+1 == combination.Yellow || combination.Yellow+1 == combination.Purple {
					if !verifiers[1].TestCombination(combination) {
						t.Errorf("Criteria24 verifier 1 failed with combination %v", combination)
					}
				} else {
					if !verifiers[2].TestCombination(combination) {
						t.Errorf("Criteria24 verifier 2 failed with combination %v", combination)
					}
				}
			}
		}
	}
}

// criteria 25 checks if there is no sequence of asecnding or descending numbers,
// if there are 2 numbers in a sequence of ascending or descending numbers,
// or if there are 3 numbers in a sequence of ascending or descending numbers
func TestCriteria25(t *testing.T) {
	criteria := criteriaList[25]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				if (combination.Blue+1 == combination.Yellow && combination.Yellow+1 == combination.Purple) ||
					(combination.Blue-1 == combination.Yellow && combination.Yellow-1 == combination.Purple) {
					if !verifiers[2].TestCombination(combination) {
						t.Errorf("Criteria25 verifier 0 failed with combination %v", combination)
					}
				} else if (combination.Blue+1 == combination.Yellow || combination.Yellow+1 == combination.Purple) ||
					(combination.Blue-1 == combination.Yellow || combination.Yellow-1 == combination.Purple) {
					if !verifiers[1].TestCombination(combination) {
						t.Errorf("Criteria25 verifier 1 failed with combination %v", combination)
					}
				} else {
					if !verifiers[0].TestCombination(combination) {
						t.Errorf("Criteria25 verifier 2 failed with combination %v", combination)
					}
				}
			}
		}
	}
}

// criteria 26 checks that a specific color is less than 3
func TestCriteria26(t *testing.T) {
	criteria := criteriaList[26]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				if combination.Blue < 3 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria26 verifier 0 failed with combination %v", combination)
				}
				if combination.Blue >= 3 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria26 verifier 0 failed with combination %v", combination)
				}

				if combination.Yellow < 3 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria26 verifier 1 failed with combination %v", combination)
				}
				if combination.Yellow >= 3 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria26 verifier 1 failed with combination %v", combination)
				}

				if combination.Purple < 3 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria26 verifier 2 failed with combination %v", combination)
				}
				if combination.Purple >= 3 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria26 verifier 2 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 27 checks that a specific color is less than 4
func TestCriteria27(t *testing.T) {
	criteria := criteriaList[27]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				if combination.Blue < 4 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria27 verifier 0 failed with combination %v", combination)
				}
				if combination.Blue >= 4 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria27 verifier 0 failed with combination %v", combination)
				}

				if combination.Yellow < 4 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria27 verifier 1 failed with combination %v", combination)
				}
				if combination.Yellow >= 4 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria27 verifier 1 failed with combination %v", combination)
				}

				if combination.Purple < 4 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria27 verifier 2 failed with combination %v", combination)
				}
				if combination.Purple >= 4 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria27 verifier 2 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 28 checks that a specific color is equal to 1
func TestCriteria28(t *testing.T) {
	criteria := criteriaList[28]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				if combination.Blue == 1 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria28 verifier 0 failed with combination %v", combination)
				}
				if combination.Blue != 1 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria28 verifier 0 failed with combination %v", combination)
				}

				if combination.Yellow == 1 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria28 verifier 1 failed with combination %v", combination)
				}
				if combination.Yellow != 1 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria28 verifier 1 failed with combination %v", combination)
				}

				if combination.Purple == 1 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria28 verifier 2 failed with combination %v", combination)
				}
				if combination.Purple != 1 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria28 verifier 2 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 29 checks that a specific color is equal to 3
func TestCriteria29(t *testing.T) {
	criteria := criteriaList[29]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				if combination.Blue == 3 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria29 verifier 0 failed with combination %v", combination)
				}
				if combination.Blue != 3 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria29 verifier 0 failed with combination %v", combination)
				}

				if combination.Yellow == 3 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria29 verifier 1 failed with combination %v", combination)
				}
				if combination.Yellow != 3 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria29 verifier 1 failed with combination %v", combination)
				}

				if combination.Purple == 3 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria29 verifier 2 failed with combination %v", combination)
				}
				if combination.Purple != 3 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria29 verifier 2 failed with combination %v", combination)
				}
			}
		}
	}
}

// critera 30 checks that a specific color is equal to 4
func TestCriteria30(t *testing.T) {
	criteria := criteriaList[30]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				if combination.Blue == 4 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria30 verifier 0 failed with combination %v", combination)
				}
				if combination.Blue != 4 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria30 verifier 0 failed with combination %v", combination)
				}

				if combination.Yellow == 4 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria30 verifier 1 failed with combination %v", combination)
				}
				if combination.Yellow != 4 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria30 verifier 1 failed with combination %v", combination)
				}

				if combination.Purple == 4 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria30 verifier 2 failed with combination %v", combination)
				}
				if combination.Purple != 4 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria30 verifier 2 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 31 checks that a specific color is greater than 1
func TestCriteria31(t *testing.T) {
	criteria := criteriaList[31]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}
				if combination.Blue > 1 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria31 verifier 0 failed with combination %v", combination)
				}
				if combination.Blue <= 1 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria31 verifier 0 failed with combination %v", combination)
				}

				if combination.Yellow > 1 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria31 verifier 1 failed with combination %v", combination)
				}
				if combination.Yellow <= 1 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria31 verifier 1 failed with combination %v", combination)
				}

				if combination.Purple > 1 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria31 verifier 2 failed with combination %v", combination)
				}
				if combination.Purple <= 1 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria31 verifier 2 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 32 checks that a specific color is greater than 3
func TestCriteria32(t *testing.T) {
	criteria := criteriaList[32]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combination := NumberCombination{blue, yellow, purple}

				if combination.Blue > 3 && !verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria32 verifier 0 failed with combination %v", combination)
				}
				if combination.Blue <= 3 && verifiers[0].TestCombination(combination) {
					t.Errorf("Criteria32 verifier 0 failed with combination %v", combination)
				}

				if combination.Yellow > 3 && !verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria32 verifier 1 failed with combination %v", combination)
				}
				if combination.Yellow <= 3 && verifiers[1].TestCombination(combination) {
					t.Errorf("Criteria32 verifier 1 failed with combination %v", combination)
				}

				if combination.Purple > 3 && !verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria32 verifier 2 failed with combination %v", combination)
				}
				if combination.Purple <= 3 && verifiers[2].TestCombination(combination) {
					t.Errorf("Criteria32 verifier 2 failed with combination %v", combination)
				}
			}
		}
	}
}

// criteria 33 checks that a specific color is even or odd
func TestCriteria33(t *testing.T) {
	criteria := criteriaList[33]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 6 {
		t.Errorf("Expected 6 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {

				if blue%2 == 0 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria33 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow%2 == 0 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria33 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple%2 == 0 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria33 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue%2 != 0 && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria33 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow%2 != 0 && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria33 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}

				if purple%2 != 0 && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria33 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 34 checks that a specific color is the smallest
// it can also be equal with another color
func TestCriteria34(t *testing.T) {
	criteria := criteriaList[34]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {

				if blue <= yellow && blue <= purple && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria34 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow <= blue && yellow <= purple && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria34 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple <= blue && purple <= yellow && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria34 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 35 checks that a specific color is the largest
// it can also be equal with another color
func TestCriteria35(t *testing.T) {
	criteria := criteriaList[35]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {

				if blue >= yellow && blue >= purple && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria35 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow >= blue && yellow >= purple && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria35 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple >= blue && purple >= yellow && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria35 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 36 checks that the sum of all 3 colors is a multiple of 3, 4, or 5
func TestCriteria36(t *testing.T) {
	criteria := criteriaList[36]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}

	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				sum := blue + yellow + purple
				if sum%3 == 0 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria36 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if sum%4 == 0 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria36 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if sum%5 == 0 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria36 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 37 checks that the sum of two colors equals 4
func TestCriteria37(t *testing.T) {
	criteria := criteriaList[37]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}

	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if blue+yellow == 4 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria37 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue+purple == 4 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria37 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow+purple == 4 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria37 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 38 checks that the sum of two colors equals 6
func TestCriteria38(t *testing.T) {
	criteria := criteriaList[38]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 3 {
		t.Errorf("Expected 3 verifiers, got %d", len(verifiers))
	}

	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if blue+yellow == 6 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria38 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue+purple == 6 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria38 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow+purple == 6 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria38 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 39 compares one specific color to 1
func TestCriteria39(t *testing.T) {
	criteria := criteriaList[39]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 6 {
		t.Errorf("Expected 6 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {

				if blue == 1 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria39 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow == 1 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria39 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple == 1 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria39 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue > 1 && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria39 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow > 1 && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria39 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple > 1 && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria39 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 40 compares a specific color to 3
func TestCriteria40(t *testing.T) {
	criteria := criteriaList[40]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 9 {
		t.Errorf("Expected 9 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {

				if blue < 3 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria40 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow < 3 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria40 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple < 3 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria40 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue == 3 && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria40 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow == 3 && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria40 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple == 3 && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria40 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue > 3 && !verifiers[6].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria40 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow > 3 && !verifiers[7].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria40 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple > 3 && !verifiers[8].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria40 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 41 compares a specific color to 4
func TestCriteria41(t *testing.T) {
	criteria := criteriaList[41]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 9 {
		t.Errorf("Expected 9 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {

				if blue < 4 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria41 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow < 4 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria41 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple < 4 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria41 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue == 4 && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria41 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow == 4 && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria41 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple == 4 && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria41 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue > 4 && !verifiers[6].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria41 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow > 4 && !verifiers[7].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria41 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purple > 4 && !verifiers[8].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria41 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// critiera 42 checks that a specific color is the smallest or largest
func TestCriteria42(t *testing.T) {
	criteria := criteriaList[42]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 6 {
		t.Errorf("Expected 6 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				blueSmallest := blue < yellow && blue < purple
				yellowSmallest := yellow < blue && yellow < purple
				purpleSmallest := purple < blue && purple < yellow
				blueLargest := blue > yellow && blue > purple
				yellowLargest := yellow > blue && yellow > purple
				purpleLargest := purple > blue && purple > yellow
				if blueSmallest && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria42 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellowSmallest && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria42 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purpleSmallest && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria42 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blueLargest && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria42 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellowLargest && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria42 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if purpleLargest && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria42 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 43 compares blue to another color
func TestCriteria43(t *testing.T) {
	criteria := criteriaList[43]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 6 {
		t.Errorf("Expected 6 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if blue < yellow && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria43 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue == yellow && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria43 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue > yellow && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria43 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue < purple && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria43 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue == purple && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria43 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue > purple && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria43 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 44 compares yellow to another color
func TestCriteria44(t *testing.T) {
	criteria := criteriaList[44]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 6 {
		t.Errorf("Expected 6 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if yellow < blue && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria44 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow == blue && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria44 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow > blue && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria44 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow < purple && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria44 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow == purple && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria44 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow > purple && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria44 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 45 tests how many 1's or 3's are in the combination
func TestCriteria45(t *testing.T) {
	criteria := criteriaList[45]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 6 {
		t.Errorf("Expected 6 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				ones := 0
				threes := 0
				if blue == 1 {
					ones++
				}
				if blue == 3 {
					threes++
				}
				if yellow == 1 {
					ones++
				}
				if yellow == 3 {
					threes++
				}
				if purple == 1 {
					ones++
				}
				if purple == 3 {
					threes++
				}
				if ones == 0 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria45 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if ones == 1 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria45 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if ones == 2 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria45 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if threes == 0 && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria45 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if threes == 1 && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria45 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if threes == 2 && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria45 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 46 checks how many 3's or 4's are in the combination
func TestCriteria46(t *testing.T) {
	criteria := criteriaList[46]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 6 {
		t.Errorf("Expected 6 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				threes := 0
				fours := 0
				if blue == 3 {
					threes++
				}
				if blue == 4 {
					fours++
				}
				if yellow == 3 {
					threes++
				}
				if yellow == 4 {
					fours++
				}
				if purple == 3 {
					threes++
				}
				if purple == 4 {
					fours++
				}
				if threes == 0 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria46 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if threes == 1 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria46 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if threes == 2 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria46 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if fours == 0 && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria46 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if fours == 1 && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria46 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if fours == 2 && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria46 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 47 checks how many 1's or 4's are in the combination
func TestCriteria47(t *testing.T) {
	criteria := criteriaList[47]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 6 {
		t.Errorf("Expected 6 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				ones := 0
				fours := 0
				if blue == 1 {
					ones++
				}
				if blue == 4 {
					fours++
				}
				if yellow == 1 {
					ones++
				}
				if yellow == 4 {
					fours++
				}
				if purple == 1 {
					ones++
				}
				if purple == 4 {
					fours++
				}
				if ones == 0 && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria47 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if ones == 1 && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria47 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if ones == 2 && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria47 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if fours == 0 && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria47 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if fours == 1 && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria47 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if fours == 2 && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria47 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}

// criteria 48 compares one specific color to another specific color
func TestCriteria48(t *testing.T) {
	criteria := criteriaList[48]
	verifiers := criteria.GetVerifiers()
	if len(verifiers) != 9 {
		t.Errorf("Expected 9 verifiers, got %d", len(verifiers))
	}
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				if blue < yellow && !verifiers[0].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria48 verifier 0 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue < purple && !verifiers[1].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria48 verifier 1 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow < purple && !verifiers[2].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria48 verifier 2 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue == yellow && !verifiers[3].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria48 verifier 3 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue == purple && !verifiers[4].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria48 verifier 4 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow == purple && !verifiers[5].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria48 verifier 5 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue > yellow && !verifiers[6].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria48 verifier 6 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if blue > purple && !verifiers[7].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria48 verifier 7 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
				if yellow > purple && !verifiers[8].TestCombination(NumberCombination{blue, yellow, purple}) {
					t.Errorf("Criteria48 verifier 8 failed with combination %v", NumberCombination{blue, yellow, purple})
				}
			}
		}
	}
}
