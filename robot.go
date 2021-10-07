package main

type Facing int

const (
	// Enum for compass directions
	North Facing = iota
	East
	South
	West

	// Default tabletop dimensions
	MaxWidth  int = 5
	MaxLength int = 5
)

// String representation of facing enum
func (f Facing) String() string {
	return [...]string{"North", "East", "South", "West"}[f]
}

type Movable interface {
	moveForward()
	turnLeft()
	turnRight()
	report() string
}

type Tabletop struct {
	width  int
	length int
}

type Robot struct {
	xPos     int
	yPos     int
	Facing   Facing
	tabletop Tabletop
}

// func (r Robot) step() error {

// }

func (r *Robot) turnLeft() {
	if r.Facing == North {
		r.Facing = West
	} else {
		r.Facing--
	}
}

func (r *Robot) turnRight() {
	if r.Facing == West {
		r.Facing = North
	} else {
		r.Facing++
	}
}
