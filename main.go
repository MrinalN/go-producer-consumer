package main

import (
	"fmt"
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

func makePizza(pizzaNumber int) *PizzaOrder {
	// increment pizza number by 1
	pizzaNumber++

	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order number #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		successs := false

		// Select State conditionals
		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza number %d. It will take %d seconds.\n", pizzaNumber, delay)

		// delay a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingrediants for pizza #%d!", rune(pizzaNumber))
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The chef quit before making for pizza #%d!", rune(pizzaNumber))
		} else {
			successs = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			successs:    successs,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// run forever or until we recieve quit notification
	for {
		// when running, try to make pizzas (in background)
		currentPizza := makePizza(i)

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
