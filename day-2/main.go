package main

import (
	"fmt"
	"os"
	"strings"
)

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	panic("Value not found in index")
}

const drawValue = 3
const winValue = 6
const lostValue = 0

// const rockValue = 1
// const paperValue = 2
// const scissorsValue = 3

const lost = "X"
const draw = "Y"
const win = "Z"

func guessWhatToPlay(opponent int, outcome string) int {
	if outcome == lost {
		if opponent == 1 {
			return 3
		} else {
			return opponent - 1
		}
	}

	if outcome == draw {
		return opponent
	}

	if outcome == win {
		if opponent == 3 {
			return 1
		} else {
			return opponent + 1
		}
	}

	panic("Unguessable next play")
}

func main() {
	data, _ := os.ReadFile("./input")

	score := 0
	for _, rawRound := range strings.Split(string(data), "\n") {
		if rawRound == "" {
			continue
		}

		round := strings.Split(rawRound, " ")

		opponentValues := []string{"A", "B", "C"}
		opponent := round[0]
		opponentValue := indexOf(opponent, opponentValues) + 1

		// meValues := []string{"X", "Y", "Z"}
		// me := round[1]
		// meValue := indexOf(me, meValues) + 1
		// roundStatuses := []string{"X", "Y", "Z"}
		roundStatus := round[1]
		// roundStatusValue := indexOf(roundStatus, roundStatuses) + 1

		meValue := guessWhatToPlay(opponentValue, roundStatus)

		switch roundStatus {
		case lost:
			score += meValue + lostValue
		case draw:
			score += meValue + drawValue
		case win:
			score += meValue + winValue
		}
	}

	fmt.Println(score)
}
