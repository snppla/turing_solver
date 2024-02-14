package solver

import (
	"sort"
	"turing_solver/criteria"

	"golang.org/x/tools/container/intsets"
)

type History struct {
	VerifierIndexes []int
	Combination     criteria.NumberCombination
}

func generateAllCombinations() []criteria.NumberCombination {
	combinations := make([]criteria.NumberCombination, 0)
	for blue := 1; blue <= 5; blue++ {
		for yellow := 1; yellow <= 5; yellow++ {
			for purple := 1; purple <= 5; purple++ {
				combinations = append(combinations, criteria.NumberCombination{Blue: blue, Yellow: yellow, Purple: purple})
			}
		}
	}
	return combinations
}

func getPossibleCombinations(remainingCombinations []criteria.NumberCombination, criteriaList []criteria.Criteria, history History) []History {

	if len(remainingCombinations) == 0 {
		return make([]History, 0)
	}
	if len(criteriaList) == 0 {
		if len(remainingCombinations) != 1 {
			return make([]History, 0)
		}
		history.Combination = remainingCombinations[0]
		return []History{history}
	}

	filteredCombinations := make([]History, 0)
	for index, verifier := range criteriaList[0].GetVerifiers() {
		verifiedCombinations := make([]criteria.NumberCombination, 0)
		for _, combination := range remainingCombinations {
			if verifier.TestCombination(combination) {
				verifiedCombinations = append(verifiedCombinations, combination)
			}
		}

		if len(verifiedCombinations) != 0 && len(verifiedCombinations) < len(remainingCombinations) {
			newHistory := History{VerifierIndexes: append(history.VerifierIndexes, index)}
			filteredCombinations = append(filteredCombinations, getPossibleCombinations(verifiedCombinations, criteriaList[1:], newHistory)...)
		}

	}

	return filteredCombinations
}

func permutations[T interface{}](arr []T) [][]T {
	var helper func([]T, int)
	res := [][]T{}

	helper = func(arr []T, n int) {
		if n == 1 {
			tmp := make([]T, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func GetPossiblePaths(criteriaList []criteria.Criteria, verifiersToInclude []*intsets.Sparse, verifiersToExclude []*intsets.Sparse) []History {
	var possibleCombinations map[string]History

	criteraPermutations := permutations(criteriaList)
	for i, permutation := range criteraPermutations {
		calculatedPaths := getPossibleCombinations(generateAllCombinations(), permutation, History{})
		calculatedSet := make(map[string]History)
		for _, path := range calculatedPaths {
			calculatedSet[path.Combination.ToString()] = path
		}
		if i == 0 {
			possibleCombinations = calculatedSet
		} else {
			for key := range possibleCombinations {
				if _, ok := calculatedSet[key]; !ok {
					delete(possibleCombinations, key)
				}
			}
		}
	}

	returnValue := make([]History, 0)
NEXT_COMBINATION:
	for _, value := range possibleCombinations {
		for verifierIndex := range value.VerifierIndexes {
			if verifiersToExclude[verifierIndex] != nil && verifiersToExclude[verifierIndex].Has(value.VerifierIndexes[verifierIndex]) {
				continue NEXT_COMBINATION
			}
			if verifiersToInclude[verifierIndex] != nil && !verifiersToInclude[verifierIndex].Has(value.VerifierIndexes[verifierIndex]) {
				continue NEXT_COMBINATION
			}
		}
		returnValue = append(returnValue, value)
	}

	sort.Slice(returnValue, func(i, j int) bool {
		numbersI := []int{returnValue[i].Combination.Blue, returnValue[i].Combination.Yellow, returnValue[i].Combination.Purple}
		numbersJ := []int{returnValue[j].Combination.Blue, returnValue[j].Combination.Yellow, returnValue[j].Combination.Purple}
		for k := 0; k < 3; k++ {
			if numbersI[k] < numbersJ[k] {
				return true
			}
			if numbersI[k] > numbersJ[k] {
				return false
			}
		}
		return false
	})

	// look through the return value and remove any duplicates of the VerifierIndexes
	for i := 0; i < len(returnValue)-1; i++ {
		duplicateFound := false
		for j := i + 1; j < len(returnValue); j++ {
			for k := 0; k < len(criteriaList); k++ {
				if returnValue[i].VerifierIndexes[k] != returnValue[j].VerifierIndexes[k] {
					break
				}
				if k == len(criteriaList)-1 {
					returnValue = append(returnValue[:j], returnValue[j+1:]...)
					j--
					duplicateFound = true
				}
			}
		}
		if duplicateFound {
			returnValue = append(returnValue[:i], returnValue[i+1:]...)
			i--

		}
	}

	return returnValue
}
