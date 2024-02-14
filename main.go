package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"turing_solver/criteria"
	"turing_solver/solver"

	"github.com/akamensky/argparse"
	"golang.org/x/tools/container/intsets"
)

func main() {
	parser := argparse.NewParser("Turing Solver", "Solves the Turing game")

	criteriaStringList := parser.String("c", "criteria", &argparse.Options{Required: true, Help: "List of criteria cards, comma separated"})
	verifiersToIncludeStringList := parser.String("i", "verifiers", &argparse.Options{Required: false, Help: "List of verifiers to include, comma separated. Example a2, e3", Default: ""})
	verifiersToExcludeStringList := parser.String("x", "exclude", &argparse.Options{Required: false, Help: "List of verifiers to exclude, comma separated. Example a2, e3", Default: ""})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		panic(err)
	}

	criteraCards := criteria.GetCriteriaList()

	cardsToPlay := make([]criteria.Criteria, 0)
	for _, card := range strings.Split(*criteriaStringList, ",") {
		cardIndex, err := strconv.Atoi(card)
		if err != nil {
			fmt.Printf("Invalid card index: %d\n", cardIndex)
			os.Exit(1)
		}
		if cardIndex < 1 || cardIndex >= len(criteraCards) {
			fmt.Printf("Invalid card index: %d\n", cardIndex)
			os.Exit(1)
		}
		cardsToPlay = append(cardsToPlay, criteraCards[cardIndex])
	}

	verifiersToInclude := make([]*intsets.Sparse, len(cardsToPlay))
	for _, verifier := range strings.Split(*verifiersToIncludeStringList, ",") {
		if verifier == "" {
			continue
		}
		critieriaIndex := int(strings.ToLower(verifier)[0] - 'a')
		verifierIndex, err := strconv.Atoi(verifier[1:])
		if err != nil {
			fmt.Printf("Invalid verifier index: %d\n", verifierIndex)
			os.Exit(1)
		}
		if verifiersToInclude[critieriaIndex] == nil {
			verifiersToInclude[critieriaIndex] = &intsets.Sparse{}
		}
		verifiersToInclude[critieriaIndex].Insert(verifierIndex)
	}

	verifiersToExclude := make([]*intsets.Sparse, len(cardsToPlay))
	for _, verifier := range strings.Split(*verifiersToExcludeStringList, ",") {
		if verifier == "" {
			continue
		}
		critieriaIndex := int(strings.ToLower(verifier)[0] - 'a')
		verifierIndex, err := strconv.Atoi(verifier[1:])
		if err != nil {
			fmt.Printf("Invalid verifier index: %d\n", verifierIndex)
			os.Exit(1)
		}
		if verifiersToExclude[critieriaIndex] == nil {
			verifiersToExclude[critieriaIndex] = &intsets.Sparse{}
		}
		verifiersToExclude[critieriaIndex].Insert(verifierIndex)
	}

	possiblePaths := solver.GetPossiblePaths(cardsToPlay, verifiersToInclude, verifiersToExclude)

	for _, paths := range possiblePaths {
		fmt.Printf("%v with verifiers %v\n", paths.Combination, paths.VerifierIndexes)
	}

	// map which criteria cards had multiple verifiers used
	verifiersUsed := map[int]*intsets.Sparse{} // Change the type to map[int]*intsets.Sparse
	for _, path := range possiblePaths {
		for criteriaIndex, verifierIndex := range path.VerifierIndexes {
			if _, ok := verifiersUsed[criteriaIndex]; !ok {
				verifiersUsed[criteriaIndex] = &intsets.Sparse{}
			}
			verifiersUsed[criteriaIndex].Insert(verifierIndex)
		}
	}

	// print out the criteria cards that had multiple verifiers used

	for criteriaIndex := 0; criteriaIndex < len(verifiersUsed); criteriaIndex++ {
		verifierSet := verifiersUsed[criteriaIndex]
		if verifierSet.Len() > 1 {
			fmt.Printf("Criteria card %c had multiple verifiers used: %v\n", 'A'+criteriaIndex, verifierSet)
		}
	}

}
