package main

import (
	"errors"
	"fmt"
)

type Facing int

const (
	// Enum for compass directions
	North Facing = iota
	East
	South
	West

	// Default tabletop dimensions
	XMin int = 0
	YMin int = 0
	XMax int = 4
	YMax int = 4
)

// String representation of facing enum
func (f Facing) String() string {
	return [...]string{"North", "East", "South", "West"}[f]
}

// type Movable interface {
// 	moveForward()
// 	turnLeft()
// 	turnRight()
// 	report() string
// }

type Tabletop struct {
	width  int
	length int
	robot  *Robot
}

type Robot struct {
	xPos   int
	yPos   int
	Facing Facing
}

// func (r Robot) step() error {

// }

// Places a robot on the table given valid coordinates and facing
func (t *Tabletop) PlaceRobot(xPos int, yPos int, facing Facing) error {
	if xPos < XMin || yPos < YMin || xPos > XMax || yPos > YMax {
		return errors.New("invalid robot coordinates")
	} else if t.robot != nil {
		return errors.New("robot has already been placed")
	}

	t.robot = &Robot{xPos: xPos, yPos: yPos, Facing: facing}
	return nil
}

// Rotate the robot 90 degrees counter-clockwise
func (t *Tabletop) TurnRobotLeft() {
	if t.robot.Facing == North {
		t.robot.Facing = West
	} else {
		t.robot.Facing--
	}
}

// Rotate the robot 90 degrees clockwise
func (t *Tabletop) TurnRobotRight() {
	if t.robot.Facing == West {
		t.robot.Facing = North
	} else {
		t.robot.Facing++
	}
}

// Move the robot one unit in the direction it is facing
func (t *Tabletop) MoveRobot() {
	if t.robot.Facing == North && t.robot.yPos != t.length {
		t.robot.yPos++
	} else if t.robot.Facing == East && t.robot.xPos != t.width {
		t.robot.xPos++
	} else if t.robot.Facing == South && t.robot.yPos != YMin {
		t.robot.yPos--
	} else if t.robot.Facing == West && t.robot.xPos != XMin {
		t.robot.xPos--
	}
}

// Report the location of the robot to stdout
func (t *Tabletop) Report() {
	fmt.Printf("%d, %d, %s\n", t.robot.xPos, t.robot.yPos, t.robot.Facing.String())
}
