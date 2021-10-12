package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tabletop := Tabletop{xMin: XMin, yMin: YMin, xMax: XMax, yMax: YMax, robot: nil}

	// Take input until a valid PLACE command is given
	err := readPlaceCommand(reader, &tabletop)
	for err != nil {
		err = readPlaceCommand(reader, &tabletop)
	}

	// Read MOVE and REPORT commands
	for {
		readRobotInput(reader, &tabletop)
	}
}

// Read the initial PLACE command and position robot on table
func readPlaceCommand(r *bufio.Reader, t *Tabletop) error {
	var (
		command   string
		xPos      int
		yPos      int
		facingStr string
		facing    Facing
	)

	// Read expected command from stdin
	n, err := fmt.Fscanf(r, "%s %d,%d,%s", &command, &xPos, &yPos, &facingStr)
	if err != nil || n != 4 {
		return errors.New("invalid input")
	}

	// Convert facing string into enum representation
	if facingStr == North.String() {
		facing = North
	} else if facingStr == East.String() {
		facing = East
	} else if facingStr == South.String() {
		facing = South
	} else if facingStr == West.String() {
		facing = West
	} else {
		return errors.New("invalid facing")
	}

	if command != "PLACE" {
		return errors.New("PLACE command expected")
	}

	err = t.PlaceRobot(xPos, yPos, facing)
	if err != nil {
		return err
	}

	return nil
}

// Read LEFT, RIGHT, MOVE and REPORT commands
func readRobotInput(r *bufio.Reader, t *Tabletop) {
	var command string

	fmt.Fscanf(r, "%s", &command)
	command = strings.Trim(command, " ")

	if command == "MOVE" {
		t.MoveRobot()
	} else if command == "LEFT" {
		t.TurnRobotLeft()
	} else if command == "RIGHT" {
		t.TurnRobotRight()
	} else if command == "REPORT" {
		t.Report()
	}
}
