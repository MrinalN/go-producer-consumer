package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// channels!
type Producer struct {
	data chan PizzaOrder
	quit chan chan error // channel of channels of errors!
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	successs    bool
}

// Close is simply a method of closing the channel when we are done with it (i.e.
// something is pushed to the quit channel)
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch // <- operator represents the idea of passing a value from channel to a reference. Dequeing from queue and assigning value to target element.
	return <-ch
}

// makePizza attempts to make a pizza. We generate a random number from 1-12,
// and put in two cases where we can't make the pizza in time. Otherwise,
// we make the pizza without issue. To make things interesting, each pizza
// will take a different length of time to produce (some pizzas are harder than others).
func makePizza(pizzaNumber int) *PizzaOrder {
	// increment pizza number by 1
	pizzaNumber++

	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order number #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		successs := false

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

// pizzeria is a goroutine that runs in the background and
// calls makePizza to try to make one order each time it iterates through
// the for loop. It executes until it receives something on the quit
// channel. The quit channel does not receive anything until the consumer
// sends it (when the number of orders is greater than or equal to the
// constant NumberOfPizzas).
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// run forever or until we recieve quit notification
	for {
		// when running, try to make pizzas (in background)
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber

			// Select State - only usable in channels - like Switch statements
			select {
			// tried to make pizza (sent something to data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:

				// close channels!
				close(pizzaMaker.data)
				close(quitChan)
				return // get the fuck out of this nested loop
			}
		}

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
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.successs {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Customer delivery unsuccessful")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()

			if err != nil {
				color.Red("*** Error closing the channel", err)
			}
		}
	}

	// print out end message
}
