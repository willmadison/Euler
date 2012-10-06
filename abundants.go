package main

import (
  "fmt"
  "os"
)

type Candidate struct {
  Number int
  ProperDivisors []int
}

type AbundantPair struct {
  first,second Candidate
}

func (c Candidate) calculateDivisorSum() int {
  sum := 0

  for _, divisor := range(c.ProperDivisors) {
    sum += divisor
  }

  return sum
}

func (c Candidate) isAbundant() bool {
  divisorSum := c.calculateDivisorSum()

  return divisorSum > c.Number
}

func (a AbundantPair) sum() int {
  return a.first.Number + a.second.Number
}

func newCandidate(number int) Candidate {
  candidate := Candidate{Number: number, ProperDivisors : []int{}}

  for i := 1; i < number; i++ {
    if number % i == 0 {
      candidate.ProperDivisors = append(candidate.ProperDivisors, i)
    }
  }

  return candidate
}

var pairwiseAbundantsBySum map[int]AbundantPair = map[int]AbundantPair{}

func main() {

  abundantCandidates := []Candidate{}

  for number := 1; number < 28123; number++ {
    candidate := newCandidate(number)

    if candidate.isAbundant() {
      abundantCandidates = append(abundantCandidates, candidate)
    }
  }
  
  for _,firstAbundant := range(abundantCandidates) {
    for _,secondAbundant := range(abundantCandidates) {
      pairwiseAbundant := AbundantPair{firstAbundant, secondAbundant}

      pairwiseAbundantsBySum[pairwiseAbundant.sum()] = pairwiseAbundant
    }
  }

  sum := 0

  for number := 1; number < 28123; number++ {
    _,present := pairwiseAbundantsBySum[number]

    if !present {
      sum += number
    }
  }

  fmt.Fprintf(os.Stdout, "The sum is: %v\n", sum)
}
