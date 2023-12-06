package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./day05/input")
	if err != nil {
		fmt.Println(err)
		return
	}

	almanac := parseAlmanac(string(data))
	fmt.Println("Part 1:", partOne(almanac))
	fmt.Println("Part 2:", partTwo(almanac))
}

type Almanac struct {
	Seeds                 []int
	SeedToSoil            *Map
	SoilToFertilizer      *Map
	FertilizerToWater     *Map
	WaterToLight          *Map
	LightToTemperature    *Map
	TemperatureToHumidity *Map
	HumidityToLocation    *Map
}

func (a *Almanac) BulkTranslate(seed int, l int) int {
	seed = a.SeedToSoil.BulkTranslate(seed, l)
	seed = a.SoilToFertilizer.BulkTranslate(seed, l)
	seed = a.FertilizerToWater.BulkTranslate(seed, l)
	seed = a.WaterToLight.BulkTranslate(seed, l)
	seed = a.LightToTemperature.BulkTranslate(seed, l)
	seed = a.TemperatureToHumidity.BulkTranslate(seed, l)
	seed = a.HumidityToLocation.BulkTranslate(seed, l)

	return seed
}

func (a *Almanac) Translate(seed int) int {
	seed = a.SeedToSoil.Translate(seed)
	seed = a.SoilToFertilizer.Translate(seed)
	seed = a.FertilizerToWater.Translate(seed)
	seed = a.WaterToLight.Translate(seed)
	seed = a.LightToTemperature.Translate(seed)
	seed = a.TemperatureToHumidity.Translate(seed)
	seed = a.HumidityToLocation.Translate(seed)

	return seed
}

type Map struct {
	SourceStart      []int
	DestinationStart []int
	Length           []int
}

func (m *Map) BulkTranslate(seed int, l int) int {
	for i := 0; i < len(m.Length); i++ {
		src := m.SourceStart[i]
		dst := m.DestinationStart[i]
		mapLen := m.Length[i]

		// need to check what amount of the upper range is changed or not
		// AKA write a better algo and think about this more
		if seed >= src && seed <= src+mapLen || seed+l >= src && seed+l <= src+mapLen {
			offset := seed - src
			return dst + offset
		}
	}

	return seed
}

func (m *Map) Translate(seed int) int {
	for i := 0; i < len(m.Length); i++ {
		src := m.SourceStart[i]
		dst := m.DestinationStart[i]
		mapLen := m.Length[i]

		if seed >= src && seed <= src+mapLen {
			offset := seed - src
			return dst + offset
		}
	}

	return seed
}

func parseAlmanac(data string) *Almanac {
	a := &Almanac{Seeds: make([]int, 0)}

	var currentMap *Map
	for i, line := range strings.Split(data, "\n") {
		if i == 0 {
			seedsString := strings.ReplaceAll(line, "seeds: ", "")
			for _, seed := range strings.Split(seedsString, " ") {
				s, _ := strconv.Atoi(seed)
				a.Seeds = append(a.Seeds, s)
			}
			continue
		}

		if line == "" {
			currentMap = nil
		} else if strings.Contains(line, "map:") {
			currentMap = &Map{
				SourceStart:      make([]int, 0),
				DestinationStart: make([]int, 0),
				Length:           make([]int, 0),
			}

			switch line {
			case "seed-to-soil map:":
				a.SeedToSoil = currentMap
			case "soil-to-fertilizer map:":
				a.SoilToFertilizer = currentMap
			case "fertilizer-to-water map:":
				a.FertilizerToWater = currentMap
			case "water-to-light map:":
				a.WaterToLight = currentMap
			case "light-to-temperature map:":
				a.LightToTemperature = currentMap
			case "temperature-to-humidity map:":
				a.TemperatureToHumidity = currentMap
			case "humidity-to-location map:":
				a.HumidityToLocation = currentMap
			default:
				fmt.Println("This cannot happen", line)
				os.Exit(1)
			}
		} else {
			nums := strings.Fields(line)
			dst, _ := strconv.Atoi(nums[0])
			src, _ := strconv.Atoi(nums[1])
			len, _ := strconv.Atoi(nums[2])

			currentMap.SourceStart = append(currentMap.SourceStart, src)
			currentMap.DestinationStart = append(currentMap.DestinationStart, dst)
			currentMap.Length = append(currentMap.Length, len)
		}
	}

	return a
}

func partOne(almanac *Almanac) int {
	var lowest int
	for _, seed := range almanac.Seeds {
		s := almanac.Translate(seed)
		if lowest == 0 || s < lowest {
			lowest = s
		}
	}

	return lowest
}

func partTwo(almanac *Almanac) int {
	var lowest int
	for i := 0; i < len(almanac.Seeds)/2; i += 2 {
		start := almanac.Seeds[i]
		len := almanac.Seeds[i+1]
		// probably a faster way to check all of these?
		s := almanac.BulkTranslate(start, len)
		if lowest == 0 || s < lowest {
			lowest = s
		}
	}

	return lowest
}
