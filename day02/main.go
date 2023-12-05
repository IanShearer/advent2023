package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./day02/input")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := strings.Split(string(data), "\n")
	fmt.Println("Part 1:", partOne(rows))
	fmt.Println("Part 2:", partTwo(rows))
}

type Round struct {
	Green int
	Blue  int
	Red   int
}

type Game struct {
	Rounds []Round
	ID     int
}

var numberRegex = regexp.MustCompile(`(\d+)`)

func parseNumberOutOfColor(color string) int {
	numberString := numberRegex.FindString(color)
	num, err := strconv.Atoi(numberString)
	if err != nil {
		fmt.Println("failed to convert number string", err)
		os.Exit(1)
	}

	return num
}

func parseGame(game string) Game {
	var g Game

	labelAndGame := strings.Split(game, ":")
	stringRounds := strings.Split(labelAndGame[1], ";")
	for _, r := range stringRounds {
		var round Round
		colors := strings.Split(r, ",")
		for _, color := range colors {
			if strings.Contains(color, "green") {
				round.Green = parseNumberOutOfColor(color)
			} else if strings.Contains(color, "blue") {
				round.Blue = parseNumberOutOfColor(color)
			} else if strings.Contains(color, "red") {
				round.Red = parseNumberOutOfColor(color)
			} else {
				fmt.Println("this should never happen, color:", color)
				os.Exit(1)
			}
		}

		g.Rounds = append(g.Rounds, round)
	}

	return g
}

func partOne(games []string) int {
	var sum int

	for i, game := range games {
		g := parseGame(game)
		var failed bool
		for _, round := range g.Rounds {
			if round.Red > 12 {
				failed = true
			}

			if round.Green > 13 {
				failed = true
			}

			if round.Blue > 14 {
				failed = true
			}
		}

		if !failed {
			sum += i + 1
		}
	}

	return sum
}

func partTwo(games []string) int {
	var sum int

	for _, game := range games {
		g := parseGame(game)
		var red int
		var green int
		var blue int
		for _, round := range g.Rounds {
			if round.Red > red {
				red = round.Red
			}

			if round.Green > green {
				green = round.Green
			}

			if round.Blue > blue {
				blue = round.Blue
			}
		}
		sum += (red * green * blue)
	}

	return sum
}
