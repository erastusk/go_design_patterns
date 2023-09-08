package main

import "fmt"

func main() {
	housebuilder := gethouseBuilder()
	h := housebuilder.AddDoor("wood").AddFloors(3).AddWindow("glass")
	fmt.Printf("%+v", h)
	// OR
	housebuilder_1 := housebuilder.buildHouse()
	fmt.Printf("%+v", housebuilder_1)
}
