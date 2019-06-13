package main

type candidate struct {
	number         int
	properDivisors []int
}

func (c candidate) calculateDivisorSum() int {
	sum := 0

	for _, divisor := range c.properDivisors {
		sum += divisor
	}

	return sum
}

func newCandidate(number int) candidate {
	candidate := candidate{number: number, properDivisors: []int{}}

	for i := 1; i < number; i++ {
		if number%i == 0 {
			candidate.properDivisors = append(candidate.properDivisors, i)
		}
	}

	return candidate
}

var candidatesBySum = map[int][]candidate{}
var amicables = []candidate{}

func SumAmicableNumbersThrough(n int) int {
	candidates := []candidate{}

	for number := 1; number < 10000; number++ {
		candidate := newCandidate(number)
		candidates = append(candidates, candidate)

		divisorSum := candidate.calculateDivisorSum()
		candidatesBySum[divisorSum] = append(candidatesBySum[divisorSum], candidate)
	}

	//Find all of the amicables
	for number := 1; number < n; number++ {
		candidate := newCandidate(number)
		divisorSum := candidate.calculateDivisorSum()
		complementors, present := candidatesBySum[candidate.number]

		if present {
			for _, amicableCandidate := range complementors {
				if amicableCandidate.number == divisorSum && amicableCandidate.number != candidate.number {
					//This is an amicable pair (i.e. there was a number whose sum was
					//equal to this number & it's distinct from this number)
					amicables = append(amicables, candidate)
				}
			}
		}
	}

	sum := 0

	for _, amicable := range amicables {
		sum += amicable.number
	}

	return sum
}
