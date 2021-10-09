package main

import "fmt"

func main() {
	robot := Robot{0, 0, East}
	tabletop := Tabletop{5, 5, robot}
	fmt.Println(tabletop.robot.Facing.String())
	fmt.Println((robot.Facing))
}
