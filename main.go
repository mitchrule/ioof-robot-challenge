package main

import "fmt"

type Facing int

const (
	North Facing = iota
	East
	South
	West

	MaxWidth  int = 5
	MaxLength int = 5
)

// String representation of facing enum
func (f Facing) String() string {
	return [...]string{"North", "East", "South", "West"}[f]
}

type Tabletop struct {
	width  int
	length int
	robot  Robot
}

type Robot struct {
	xpos   int
	ypos   int
	facing Facing
}

func main() {
	robot := Robot{0, 0, North}
	tabletop := Tabletop{0, 0, robot}
	fmt.Println(tabletop)
}
