package day06

type Puzzler struct {
}

func (Puzzler) Part1(input []string) string {
	times := []int{58, 99, 64, 69}
	distances := []int{478, 2232, 1019, 1071}

	margin := 1
	for i := range times {
		countWins := 0
		for maxSpeed := times[i] / 2; maxSpeed*(times[i]-maxSpeed) > distances[i]; maxSpeed-- {
			countWins++
		}

		if countWins > 0 {
			countWins = countWins * 2

			if times[i]%2 == 0 {
				countWins--
			}
			margin = margin * countWins
		}
	}
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
