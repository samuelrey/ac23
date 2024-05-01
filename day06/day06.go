package day06

type Puzzler struct {
}

func (Puzzler) Part1(input []string) string {
	times := []int{58, 99, 64, 69}
	distances := []int{478, 2232, 1019, 1071}

	/*
		for each millisecond hold the button at start (time=0; time<times[i]), speed = speed + 1 mm/s
		speed at start is always 0
		distance at start is always 0
		speed = distance / time
		time and speed have an inverse relationship, ie the more speed we have, the less time we have to race
		therefore, the most speed we could have in a race is equal to the time

		s | t | d | d > max
		0 | 7 | 0 | false
		1 | 6 | 6 | false
		2 | 5 | 10 | true
		3 | 4 | 12 | true
		4 | 3 | 12 | true
		5 | 2 | 10 | true
		6 | 1 | 6 | false
		7 | 0 | 0 | false

		15 / 2 = 7 (max speed)
		15 - 7 = 8
		s | t | d | d > max
		7 | 8 | 56 | true
		6 | 9 | 54 | true
		5 | 10 | 50 | true
		4 | 11 | 44 | true
		3 | 12 | 36 | false

		30 / 2 = 15
		30 - 15 = 15
		s | t | d | d > max
		15 | 15 | 225 | true
		14 | 16 | 224 | true
		13 | 17 | 221 | true
		12 | 18 | 216 | true
		11 | 19 | 209 | true
		10 | 20 | 200 | false

		if time % 2 == 0
			winningScenarios = (2 * winningScenarios) - 1
	*/
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
