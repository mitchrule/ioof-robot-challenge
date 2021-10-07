package main

import "fmt"

func main() {
	tabletop := Tabletop{5, 5}
	robot := Robot{0, 0, West, tabletop}
	fmt.Println(robot.Facing.String())
	robot.turnLeft()
	fmt.Println((robot.Facing))
}
