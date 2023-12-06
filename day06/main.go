package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./day06/input")
	if err != nil {
		fmt.Println(err)
		return
	}

	races := parseRacesPartOne(string(data))
	fmt.Println("Part 1:", partOne(races))
	races = parseRacesPartTwo(string(data))
	fmt.Println("Part 2:", partTwo(races))
}

type Race struct {
	TotalTime      int
	RecordDistance int
}

func (r Race) FindNumberOfWins() int {
	var differentWins int

	for i := 0; i < r.TotalTime; i++ {
		d := i * (r.TotalTime - i)
		if d > r.RecordDistance {
			differentWins++
		}
	}

	return differentWins
}

func parseRacesPartOne(data string) []Race {
	s := strings.Split(data, "\n")
	times := strings.Fields(strings.ReplaceAll(s[0], "Time:", ""))
	distances := strings.Fields(strings.ReplaceAll(s[1], "Distance:", ""))

	races := make([]Race, 0)
	for i := 0; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		races = append(races, Race{TotalTime: t, RecordDistance: d})
	}

	return races
}

func parseRacesPartTwo(data string) []Race {
	s := strings.Split(data, "\n")
	times := strings.ReplaceAll((strings.ReplaceAll(s[0], "Time:", "")), " ", "")
	distances := strings.ReplaceAll(strings.ReplaceAll(s[1], "Distance:", ""), " ", "")

	races := make([]Race, 0)
	t, _ := strconv.Atoi(times)
	d, _ := strconv.Atoi(distances)
	races = append(races, Race{TotalTime: t, RecordDistance: d})

	return races
}

func partOne(races []Race) int {
	var sum int
	for _, r := range races {
		numOfWins := r.FindNumberOfWins()
		if sum == 0 {
			sum = numOfWins
		} else {
			sum *= numOfWins
		}
	}

	return sum
}

func partTwo(races []Race) int {
	var sum int
	for _, r := range races {
		numOfWins := r.FindNumberOfWins()
		if sum == 0 {
			sum = numOfWins
		} else {
			sum *= numOfWins
		}
	}

	return sum
}
