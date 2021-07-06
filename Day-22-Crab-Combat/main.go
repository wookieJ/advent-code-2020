package main

import (
	"Day-22-Crab-Combat/common"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 22: Crab Combat  ---")
	input := common.GetInputFromFile(dataPath)

	start := time.Now()
	resultPart1 := firstPart(input)
	firstPartDuration := time.Since(start)
	fmt.Println(fmt.Sprintf("  Part 1 >> %d [after: %v]", resultPart1, firstPartDuration))

	start = time.Now()
	resultPart2 := secondPart(input)
	secondPartDuration := time.Since(start)
	fmt.Println(fmt.Sprintf("  Part 2 >> %d [after: %v]", resultPart2, secondPartDuration))
}

func firstPart(input string) int {
	lines := common.GetStringArrayFromStringInput(input, "\n\n")
	player1Deck := common.GetIntArrayFromStringInput(strings.TrimLeft(lines[0], "Player 1:\n"), "\n")
	player2Deck := common.GetIntArrayFromStringInput(strings.TrimLeft(lines[1], "Player 2:\n"), "\n")
	for true {
		player1TopCard := player1Deck[0]
		player2TopCard := player2Deck[0]
		player1Deck = player1Deck[1:]
		player2Deck = player2Deck[1:]
		if player1TopCard > player2TopCard {
			player1Deck = append(player1Deck, player1TopCard, player2TopCard)
		} else {
			player2Deck = append(player2Deck, player2TopCard, player1TopCard)
		}
		if len(player1Deck) == 0 {
			return getDeckResult(player2Deck)
		} else if len(player2Deck) == 0 {
			return getDeckResult(player1Deck)
		}
	}
	return 0
}

func getDeckResult(player1Deck []int) int {
	res := 0
	for i, card := range player1Deck {
		res += (len(player1Deck) - i) * card
	}
	return res
}

func secondPart(input string) int {
	lines := common.GetStringArrayFromStringInput(input, "\n\n")
	player1Deck := common.GetIntArrayFromStringInput(lines[0][10:], "\n")
	player2Deck := common.GetIntArrayFromStringInput(lines[1][10:], "\n")
	return recursiveGame(player1Deck, player2Deck)
}

func recursiveGame(player1Deck []int, player2Deck []int) int {
	subGame(&player1Deck, &player2Deck)
	if len(player1Deck) == 0 {
		return getDeckResult(player2Deck)
	} else if len(player2Deck) == 0 {
		return getDeckResult(player1Deck)
	}
	return 0
}

func subGame(player1Deck *[]int, player2Deck *[]int) bool {
	player1Won := true
	states := make(map[string]struct{})
	for len(*player1Deck) > 0 && len(*player2Deck) > 0 {
		state := computeStateHash(*player1Deck, *player2Deck)
		if _, ok := states[state]; ok {
			return true
		}
		states[state] = struct{}{}
		player1TopCard, player2TopCard := (*player1Deck)[0], (*player2Deck)[0]
		*player1Deck, *player2Deck = (*player1Deck)[1:], (*player2Deck)[1:]
		if len(*player1Deck) >= player1TopCard && len(*player2Deck) >= player2TopCard {
			player1Won = subGame(copyElements(player1Deck, player1TopCard), copyElements(player2Deck, player2TopCard))
		} else {
			player1Won = player1TopCard > player2TopCard
		}
		if player1Won {
			*player1Deck = append(*player1Deck, player1TopCard, player2TopCard)
		} else {
			*player2Deck = append(*player2Deck, player2TopCard, player1TopCard)
		}
	}
	if len(*player1Deck) > 0 {
		player1Won = true
	} else {
		player1Won = false
	}
	return player1Won
}

func copyElements(deck *[]int, n int) *[]int {
	var c []int
	for i := 0; i < n; i++ {
		c = append(c, (*deck)[i])
	}
	return &c
}

func computeStateHash(a, b []int) string {
	c := make([]string, len(a)+len(b)+1)
	for i, v := range a {
		c[i] = strconv.Itoa(v)
	}
	c[len(a)] = "|"
	for i, v := range b {
		c[len(a)+i+1] = strconv.Itoa(v)
	}
	return strings.Join(c, ",")
}
