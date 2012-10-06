package main

import (
  "fmt"
  "os"
)

type Candidate struct {
  Number int
  ProperDivisors []int
}

func (c Candidate) calculateDivisorSum() int {
  sum := 0

  for _, divisor := range(c.ProperDivisors) {
    sum += divisor
  }

  return sum
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

var candidatesBySum map[int][]Candidate = map[int][]Candidate{}
var amicables []Candidate = []Candidate{}

func main() {

  candidates := []Candidate{}

  for number := 1; number < 10000; number++ {
    candidate := newCandidate(number)
    candidates = append(candidates, candidate)

    divisorSum := candidate.calculateDivisorSum()
    candidatesBySum[divisorSum] = append(candidatesBySum[divisorSum], candidate)
  }

  //Find all of the amicables
  for number := 1; number < 10000; number++ {
    candidate := newCandidate(number)
    divisorSum := candidate.calculateDivisorSum()
    complementors, present := candidatesBySum[candidate.Number]

    if present {
      for _, amicableCandidate := range(complementors) {
        if amicableCandidate.Number == divisorSum && amicableCandidate.Number != candidate.Number {
          //This is an amicable pair (i.e. there was a number whose sum was 
          //equal to this number & it's distinct from this number)
          amicables = append(amicables, candidate)
        }
      }
    }
  }

  sum := 0

  for _,amicable := range(amicables) {
    sum += amicable.Number
  }

  fmt.Fprintf(os.Stdout, "The sum is: %v\n", sum)
}
