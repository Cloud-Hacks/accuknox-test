package main

import "fmt"

func main() {
	cnp := make(chan func(), 10)
	for i := 0; i < 4; i++ {
		go func() {
			for f := range cnp {
				f()
			}
		}()
	}
	cnp <- func() {
		fmt.Println("HERE1")
	}
	fmt.Println("Hello")
}

/* First, a channel cnp is created using the make function with a buffer size of 10. This channel can be used to send and receive functions.

Then, a for loop is used to create four goroutines. Each goroutine waits for functions to be received from the channel and executes them.

Inside the for loop, an anonymous function is created using a closure. This function is executed for each received function from the channel cnp.

Outside the for loop, a function is sent to the channel cnp using the <- operator. In this case, the sent function prints "HERE1" inside the go routine.

Finally, the main goroutine continues execution and prints "Hello". */
