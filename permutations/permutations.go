package permutations

import (
	"errors"
	"sort"
)

type permutation []int

func (p permutation) permute() error {
	largestIndexWithLargerSuccessor := len(p) - 2

	for p[largestIndexWithLargerSuccessor] > p[largestIndexWithLargerSuccessor+1] {
		largestIndexWithLargerSuccessor--
	}

	if largestIndexWithLargerSuccessor < 0 {
		return errors.New("No More Permutations, current permutation is the final one")
	}

	largestIndexLargerThanFoundIndex := len(p) - 1

	for p[largestIndexWithLargerSuccessor] > p[largestIndexLargerThanFoundIndex] {
		largestIndexLargerThanFoundIndex--
	}

	//Swap the values at the two indicies...
	p[largestIndexWithLargerSuccessor], p[largestIndexLargerThanFoundIndex] = p[largestIndexLargerThanFoundIndex], p[largestIndexWithLargerSuccessor]

	//Sort everything from location following Largest Index that has larger successor through the end of the slice
	sort.Ints(p[largestIndexWithLargerSuccessor+1:])

	return nil
}

var p permutation = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var permutationError error

func DetermineNthPermutation(n int) []int {
	for i := 1; i <= n && permutationError == nil; i++ {
		permutationError = p.permute()
	}

	return p
}
