package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./day01/input")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := strings.Split(string(data), "\n")
	fmt.Println("Part 1:", partOne(rows))
	fmt.Println("Part 2:", partTwo(rows))
}

func partOne(rows []string) int {

	var combined strings.Builder
	sum := 0
	for _, r := range rows {
		var firstChar rune
		var lastChar rune
		for _, c := range r {
			if c > '0' && c <= '9' {
				if firstChar == 0 {
					firstChar = c
				}

				lastChar = c
			}
		}

		if firstChar != 0 {
			combined.WriteRune(firstChar)
		}

		if lastChar != 0 {
			combined.WriteRune(lastChar)
		}

		if combined.String() != "" {
			res, err := strconv.Atoi(combined.String())
			if err != nil {
				fmt.Println(err)
				return 0
			}
			sum += res

			combined.Reset()
		}
	}

	return sum
}

func partTwo(rows []string) int {

	wordMap := map[string]rune{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	regexArr := map[string]*regexp.Regexp{
		"one":   regexp.MustCompile("one|1"),
		"two":   regexp.MustCompile("two|2"),
		"three": regexp.MustCompile("three|3"),
		"four":  regexp.MustCompile("four|4"),
		"five":  regexp.MustCompile("five|5"),
		"six":   regexp.MustCompile("six|6"),
		"seven": regexp.MustCompile("seven|7"),
		"eight": regexp.MustCompile("eight|8"),
		"nine":  regexp.MustCompile("nine|9"),
	}

	var combined strings.Builder
	sum := 0
	for _, r := range rows {
		q := make(map[int]rune)
		sortedPositionArray := make([]int, 0)
		for word, reg := range regexArr {
			m := reg.FindAllStringSubmatchIndex(r, -1)
			for _, t := range m {
				q[t[0]] = wordMap[word]
			}
		}

		for k := range q {
			sortedPositionArray = append(sortedPositionArray, k)
		}

		sort.Ints(sortedPositionArray)

		var firstChar rune
		var lastChar rune
		for _, v := range sortedPositionArray {

			rVal := q[v]

			if rVal > '0' && rVal <= '9' {
				if firstChar == 0 {
					firstChar = rVal
				}

				lastChar = rVal
			}
		}

		if firstChar != 0 {
			combined.WriteRune(firstChar)
		}

		if lastChar != 0 {
			combined.WriteRune(lastChar)
		}

		if combined.String() != "" {
			res, err := strconv.Atoi(combined.String())
			if err != nil {
				fmt.Println(err)
				return 0
			}
			sum += res

			combined.Reset()
		}

	}

	return sum
}
