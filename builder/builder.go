package main

// Builder Requirements
// 1. Product - complex struct to be generated - House
// 2. Builder - Receiver functions that create the Product - houseBuilder
// 3. Director - Directs the Build process

// https://www.geeksforgeeks.org/builder-design-pattern/

// PRODUCT
type House struct {
	window string
	door   string
	floors int
}

// BUILDER
type houseBuilder struct {
	window string
	door   string
	floors int
}

func gethouseBuilder() houseBuilder {
	return houseBuilder{}
}

func (h *houseBuilder) AddWindow(w string) *houseBuilder {
	h.window = w
	return h
}

func (h *houseBuilder) AddDoor(w string) *houseBuilder {
	h.door = w
	return h
}

func (h *houseBuilder) AddFloors(w int) *houseBuilder {
	h.floors = w
	return h
}

func (h *houseBuilder) buildHouse() House {
	return House{
		window: h.window,
		door:   h.door,
		floors: h.floors,
	}
}

// DIRECTOR
type director struct {
	builder houseBuilder
}

func newDirector(b houseBuilder) *director {
	return &director{
		builder: b,
	}
}
