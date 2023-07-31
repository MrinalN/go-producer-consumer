package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// channel!
type Producer struct {
	data chan PizzaOrder
	quit chan chan error // channel of channels of errors!
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	successs    bool
}

// func to close channel
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch // <- operator represents the idea of passing a value from channel to a reference. Dequeing from queue and assigning value to target element.
	return <-ch
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making

	// run forever or until we recieve quit notification
	for {
		// when running, try to make pizzas (in background)

		// Select State (conditionals)
	}

}

func main() {
	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	// print out start message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("------------------------------")

	// create a producer
	pizzaJob := &Producer{
		// MAKE a channel
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer

	// print out end message
}
