package abundants

type candidate struct {
  number int
  properDivisors []int
}

type abundantPair struct {
  first,second candidate
}

func (c candidate) calculateDivisorSum() int {
  sum := 0

  for _, divisor := range(c.properDivisors) {
    sum += divisor
  }

  return sum
}

func (c candidate) isAbundant() bool {
  divisorSum := c.calculateDivisorSum()

  return divisorSum > c.number
}

func (a abundantPair) sum() int {
  return a.first.number + a.second.number
}

func newCandidate(number int) candidate {
  candidate := candidate{number: number, properDivisors : []int{}}

  for i := 1; i < number; i++ {
    if number % i == 0 {
      candidate.properDivisors = append(candidate.properDivisors, i)
    }
  }

  return candidate
}

var pairwiseAbundantsBySum = map[int]abundantPair{}

func SumAbundantPairsThrough(n int) int {
  abundantCandidates := []candidate{}

  for number := 1; number < n; number++ {
    candidate := newCandidate(number)

    if candidate.isAbundant() {
      abundantCandidates = append(abundantCandidates, candidate)
    }
  }
  
  for _,firstAbundant := range(abundantCandidates) {
    for _,secondAbundant := range(abundantCandidates) {
      pairwiseAbundant := abundantPair{firstAbundant, secondAbundant}

      pairwiseAbundantsBySum[pairwiseAbundant.sum()] = pairwiseAbundant
    }
  }

  sum := 0

  for number := 1; number < n; number++ {
    _,present := pairwiseAbundantsBySum[number]

    if !present {
      sum += number
    }
  }

  return sum
}
