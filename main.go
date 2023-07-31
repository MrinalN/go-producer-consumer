package main

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// channel!
type Producer struct {
	data chan PizzaOrder
	quit chan chan error // channel of channels of errors!
}

type PizzaOrder struct {
	pizzaNumber int
	message string
	successs bool
} 

func main () {
	// seed random number generator

	// print out start message 

	// create a producer

	// run the producer in the background

	// create and run consumer

	// print out end message
}