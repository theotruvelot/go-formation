package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// create receiver function for type deck
// any variable with type deck have access to print func
// the receiver variable need to be an abbreviation of the type (here d)
// but we can make dec / de... BUT not cards... (you can but convention)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
