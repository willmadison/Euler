package main

import (
  "fmt"
  "os"
  "errors"
  "sort"
)

type Permutation []int

func (p Permutation) permute() error {
  largestIndexWithLargerSuccessor := len(p) - 2

  for p[largestIndexWithLargerSuccessor] > p[largestIndexWithLargerSuccessor + 1] {
    largestIndexWithLargerSuccessor--;
  }

  if largestIndexWithLargerSuccessor < 0 {
    return errors.New("No More Permutations, current permutation is the final one")
  }

  largestIndexLargerThanFoundIndex := len(p) - 1

  for p[largestIndexWithLargerSuccessor] > p[largestIndexLargerThanFoundIndex] {
    largestIndexLargerThanFoundIndex--;
  }

  //Swap the values at the two indicies...
  p[largestIndexWithLargerSuccessor], p[largestIndexLargerThanFoundIndex] = p[largestIndexLargerThanFoundIndex], p[largestIndexWithLargerSuccessor]

  //Sort everything from location following Largest Index that has larger successor through the end of the slice

  sort.Ints(p[largestIndexWithLargerSuccessor+1:])

  return nil
}

var permutation Permutation = []int{0,1,2,3,4,5,6,7,8,9}
var permutationError error = nil

func main() {
  for i := 1; i <= 1000000 && permutationError == nil; i++ {
    permutationError = permutation.permute()
  }

  fmt.Fprintf(os.Stdout, "The 1,000,000th permutation is: %v\n", permutation)
}
