package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./day04/input")
	if err != nil {
		fmt.Println(err)
		return
	}

	cards := parseCards(string(data))
	fmt.Println("Part 1:", partOne(cards))
	fmt.Println("Part 2:", partTwo(cards))
}

type Card struct {
	MyNumbers      map[int]struct{}
	WinningNumbers map[int]struct{}
	CardNumber     int

	numberOfWins   int
	numberOfCopies int
	checked        bool
}

func pointCalc(numOfWinningNumbers int) int {
	var num int
	for i := 0; i < numOfWinningNumbers; i++ {
		num = 2 * num
		if num == 0 {
			num = 1
		}
	}
	return num
}

func winFunc(partOne bool, n int) int {
	if partOne {
		return pointCalc(n)
	} else {
		return n
	}
}

func (c *Card) CheckWin(partOne bool) int {
	if c.checked {
		return winFunc(partOne, c.numberOfWins)
	}

	c.checked = true

	var numOfWinningNumbers int
	for num := range c.MyNumbers {
		_, ok := c.WinningNumbers[num]
		if ok {
			numOfWinningNumbers++
		}
	}

	c.numberOfWins = numOfWinningNumbers
	return winFunc(partOne, c.numberOfWins)
}

func (c *Card) CreateCopies(cardsToCopyAhead int, cards []*Card) {
	for i := c.CardNumber; i < c.CardNumber+cardsToCopyAhead; i++ {
		card := cards[i]
		card.numberOfCopies += c.numberOfCopies
	}
}

var cardRegex = regexp.MustCompile(`Card(\s+)(\d+):`)

func parseCards(s string) []*Card {
	cards := make([]*Card, 0)
	for i, c := range strings.Split(s, "\n") {
		card := &Card{
			CardNumber:     i + 1,
			MyNumbers:      make(map[int]struct{}),
			WinningNumbers: make(map[int]struct{}),
			numberOfCopies: 1,
		}
		c = cardRegex.ReplaceAllString(c, "")
		spl := strings.Split(c, "|")

		for _, num := range strings.Fields(strings.TrimSpace(spl[0])) {
			n, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			card.MyNumbers[n] = struct{}{}
		}

		for _, num := range strings.Fields(strings.TrimSpace(spl[1])) {
			n, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			card.WinningNumbers[n] = struct{}{}
		}

		cards = append(cards, card)
	}

	return cards
}

func partOne(cards []*Card) int {
	var sum int
	for _, c := range cards {
		sum += c.CheckWin(true)
	}

	return sum
}

func partTwo(cards []*Card) int {
	var sum int
	for _, c := range cards {
		wins := c.CheckWin(false)
		c.CreateCopies(wins, cards)
		sum += c.numberOfCopies
	}

	return sum
}
