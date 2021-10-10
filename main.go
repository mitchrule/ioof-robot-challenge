package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tabletop := Tabletop{width: XMax, length: YMax, robot: nil}

	err := readPlaceCommand(reader, &tabletop)
	for err != nil {
		err = readPlaceCommand(reader, &tabletop)
	}

	fmt.Println(tabletop, tabletop.robot)
}

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
