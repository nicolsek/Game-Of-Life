package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Rules! */

/*

	1. Any live cell with fewer than two live neighbors dies, as if caused by underpopulation.
	2. Any live cell with two or three live neighbors lives on to the next generation.
	3. Any live cell with more than three live neighbors dies, as if by overpopulation.
	4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

	(Rules taken from Wikipedia)

*/

// main ... Main functions that begins the simulation of life and acts as a state based loop.
func main() {
	//Creating and passing a reference to the GOL object.
	gol := makeGOL()

	//While GOL is simulating.
	for gol.isSimulating {
		//Simulating using the reference to the GOL object.
		simulate(gol)
		//Draw the GOL cells.
		draw(gol)
	}

}

// GOL ... A struct that holds the required data for simulating and managing the Game of Life.
type GOL struct {
	cells        [256][256]int
	isSimulating bool
	hasStart     bool
}

// makeGOL ... Creates the GOL object and sets the running to true.
func makeGOL() *GOL {
	gol := new(GOL)
	gol.isSimulating = true
	//Will be used to reference the start of the simulation and to allow seeding.
	gol.hasStart = false

	return gol
}

// simulate ... Begins simulating the GOL and using the object alters the values and checks the rules.
func simulate(gol *GOL) {
	if gol.hasStart {

	} else {
		seedCells(gol)
	}
}

func draw(gol *GOL) {

}

func seedCells(gol *GOL) {
	// -\(.-.)/- It's a way of seeding, don't judge me I'm 16.
	rand.Seed(time.Since(time.Now()).Nanoseconds())
	for cellsX := range gol.cells {
		for cellsY := range gol.cells {
			gol.cells[cellsX][cellsY] = rand.Intn(2)
		}
	}

	fmt.Printf("%v\n", gol.cells)
}
