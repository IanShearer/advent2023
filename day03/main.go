package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./day03/input")
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := parseGrid(string(data))
	fmt.Println("Part 1:", partOne(grid))
	fmt.Println("Part 2:", partTwo(grid))
}

type Grid struct {
	Cells  []Cell
	Width  int
	Height int
}

type Cell struct {
	X                  int
	Y                  int
	PossibleEnginePart *EnginePart
	Symbol             *Symbol
}

type Symbol struct {
	X         int
	Y         int
	Value     string
	GearRatio int

	checked bool
}

type EnginePart struct {
	StartX int
	EndX   int
	Y      int

	numberString string
	number       int
	valid        bool
	checked      bool
}

func (e *EnginePart) GetNumber() int {
	if e.number != 0 {
		return e.number
	}

	n, _ := strconv.Atoi(e.numberString)
	e.number = n
	return e.number
}

func (e *EnginePart) GetValidity(g *Grid) bool {
	if e.checked {
		return e.valid
	}

	e.checked = true
	xStart := e.StartX - 1
	xEnd := e.EndX + 1
	yStart := e.Y - 1
	yEnd := e.Y + 1

	if xStart < 0 {
		xStart = 0
	}

	if xEnd > g.Width-1 {
		xEnd = g.Width - 1
	}

	if yStart < 0 {
		yStart = 0
	}

	if yEnd > g.Height-1 {
		yEnd = g.Height - 1
	}

	// TODO: Grid placement is all messed up
	for x := xStart; x <= xEnd; x++ {
		for y := yStart; y <= yEnd; y++ {
			cell := g.GetCell(x, y)
			if cell != nil && cell.Symbol != nil {
				e.valid = true
				return e.valid
			}
		}
	}

	return e.valid
}

func (s *Symbol) CheckGearRatio(g *Grid) int {
	if s.checked {
		return s.GearRatio
	}
	s.checked = true

	if s.Value != "*" {
		return s.GearRatio
	}

	xStart := s.X - 1
	xEnd := s.X + 1
	yStart := s.Y - 1
	yEnd := s.Y + 1

	if xStart < 0 {
		xStart = 0
	}

	if xEnd > g.Width-1 {
		xEnd = g.Width - 1
	}

	if yStart < 0 {
		yStart = 0
	}

	if yEnd > g.Height-1 {
		yEnd = g.Height - 1
	}

	var firstEnginePart *EnginePart
	var secondEnginePart *EnginePart
	for x := xStart; x <= xEnd; x++ {
		for y := yStart; y <= yEnd; y++ {
			cell := g.GetCell(x, y)
			if cell != nil && cell.PossibleEnginePart != nil {
				if firstEnginePart == nil {
					firstEnginePart = cell.PossibleEnginePart
				} else if cell.PossibleEnginePart != firstEnginePart && secondEnginePart == nil {
					secondEnginePart = cell.PossibleEnginePart
				}

				if firstEnginePart != cell.PossibleEnginePart && secondEnginePart != cell.PossibleEnginePart {
					return s.GearRatio
				}
			}
		}
	}

	if firstEnginePart != nil && secondEnginePart != nil {
		s.GearRatio = firstEnginePart.GetNumber() * secondEnginePart.GetNumber()
	}

	return s.GearRatio

}

func (g *Grid) GetCell(x, y int) *Cell {
	idx := y*g.Width + x
	// fmt.Println(x, y, idx)
	return &g.Cells[idx]
}

func parseGrid(g string) *Grid {
	s := strings.Split(g, "\n")
	g = strings.ReplaceAll(g, "\n", "")

	grid := &Grid{}
	grid.Height = len(s)
	grid.Width = len(s[0])

	var currentEnginePart *EnginePart
	for i, r := range g {
		var cell Cell
		cell.X = i % (grid.Width)
		cell.Y = i / grid.Width

		if cell.X == 0 {
			currentEnginePart = nil
		}

		if r >= '0' && r <= '9' {
			if currentEnginePart == nil {
				currentEnginePart = &EnginePart{StartX: cell.X, Y: cell.Y}
			}
			currentEnginePart.numberString += string(r)
			currentEnginePart.EndX = cell.X
			cell.PossibleEnginePart = currentEnginePart
		} else if r == '.' || r == '\n' {
			currentEnginePart = nil
		} else {
			cell.Symbol = &Symbol{Value: string(r), X: cell.X, Y: cell.Y}
			currentEnginePart = nil
		}

		grid.Cells = append(grid.Cells, cell)
	}

	return grid
}

func partOne(g *Grid) int {
	var sum int
	possibleEngineParts := make(map[*EnginePart]struct{})
	for i := range g.Cells {
		p := g.Cells[i].PossibleEnginePart
		if p != nil {
			possibleEngineParts[p] = struct{}{}
		}
	}

	for e := range possibleEngineParts {
		if e.GetValidity(g) {
			sum += e.GetNumber()
		}
	}

	return sum
}

func partTwo(g *Grid) int {
	var sum int
	symbols := make(map[*Symbol]struct{})
	for i := range g.Cells {
		p := g.Cells[i].Symbol
		if p != nil {
			symbols[p] = struct{}{}
		}
	}

	for s := range symbols {
		sum += s.CheckGearRatio(g)
	}

	return sum
}
