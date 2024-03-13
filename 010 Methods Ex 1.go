// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a player's batting average. The formula is hits / atBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import "fmt"

// e a struct that represents a ball player.
// Include fields called name, atBats and hits.
type player struct{
	name string
	atBats int
	hits int
}

// Declare a method that calculates the batting average for a player.
func ( /* receiver */ p *player ) average() /* return type */ float64 {
	if p.atBats == 0 {
		return 0.0
	}
	
	return float64(p.hits) / float64(p.atBats)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	players := []player{
		{"Bill", 24, 12},
		{"Tedd", 56, 45},
	}

	// Display the batting average for each player in the slice.
	for i := range players {
		fmt.Printf("Name: %v with average %f\n", players[i].name, players[i].average())
	}
}
