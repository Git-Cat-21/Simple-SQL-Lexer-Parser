// go routines are lightweigth threads of execution in Go
// allows concurrent execution
/*package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sayHello()
	for i := 0; i < 5; i++ {
		fmt.Println("World")
		time.Sleep(100 * time.Millisecond)

	}
}*/

// channels in Go are a powerful concurrency primitive for communication and synchronization between goroutines
// provide a good way goroutines to send and receive values to and from each other to coordibate execution
/*package main

import (
	"fmt"
	"time"
)

func senData(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
		fmt.Println("Sent: ", i)
	}
	close(ch)
}

// ch <- value send value to the channel
// receivedVal := <-ch
func main() {
	ch := make(chan int)//unbuffered
	go senData(ch)
	// for num := range ch {
	// 	fmt.Println("Received ", num)
	// }
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("Received ", num)
	}
}*/

//by default channels are unbuffered they will send only if there is a corresponding receiver
//buffered chnangels accept a limited number of valuues without a corresponding receiver for those values

/*package main

import (
	"fmt"
)

func senData(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Sent: ", i)
	}
	close(ch)
}

func main() {
	ch := make(chan int, 3)
	go senData(ch)
	for num := range ch {

		fmt.Println("Data received ", num)
	}
}*/

// channel synchronization
//coordination of concurrent processes or goroutines in Go using channels.
//Blocking Operations: Sending to or receiving from a channel is a blocking operation.
//If a goroutine attempts to send to a channel and there is no receiver, it will block until a
//receiver becomes available. Similarly, if a goroutine attempts to receive from a channel and there is no sender, it will block until a sender becomes available.

// Wait Groups: The sync package in Go provides a WaitGroup type, which can be used to wait for a collection of goroutines to finish executing.
/*package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine complted its workk")
		ch <- true //send signal to the main goroutine
	}()
	fmt.Println("main gorouutine is waiting....")
	<-ch //wait for the signal from the goroutine
	fmt.Println("Main goroutine received the signal")

}*/

//channel directions
//bidirectional channel ch:=make(chan int)
//Send only chnnel ch:=make(chan <- int)
//Receive only channel ch:=make(<- chan int)
//for receive only or the send only channels if you try to try to do the other operations it results innto compile time error

/*package main

import (
	"fmt"
)

func sendval(ch chan<- string) {
	ch <- "ping"
}
func receiveVal(ch <-chan string) {
	value := <-ch
	fmt.Println(value)
}
func main() {
	ch := make(chan string)
	go sendval(ch)
	go receiveVal(ch)
	<-ch
	<-ch
	fmt.Println("DONEE")
}*/

//select
/*Go is used for communication and synchronization between goroutines.
It allows you to wait on multiple communication operations simultaneously
and proceed with the operation that can proceed first.*/
// follows the syntax similar to switch case execpt for the switch(entity)

/*package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 69
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 420
	}()
	select {
	case value := <-ch1:
		fmt.Println("Recevied from the 1st chnanel", value)
	case value := <-ch2:
		fmt.Println("Recevied from the 2nd chnanel", value)
	}
	fmt.Println("done")

}*/

/*package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		time.Sleep(10 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(10 * time.Second)
		c3 <- "three"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("terminated")
	}
}*/

// non blocking
// Non-blocking channel operations involve sending or receiving data from a channel without blocking the execution of the current goroutine.
/*package main

import "fmt"

func main() {
	messages := make(chan string)
	//signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("message received", msg)
	default:
		fmt.Println("message not received")
	}
	var msg string
	msg = "bruh"

	select {
	case messages <- msg:
		fmt.Println("message sent", msg)
	default:
		fmt.Println("messafe not sent")
	}

}*/

/*package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}

		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job ", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
	_, ok := <-jobs
	fmt.Println("recieved more jobs", ok)
}*/

//possible to close a non-emoty channel and still have the remainig values be reveived
/*package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "first"
	queue <- "second"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}*/

