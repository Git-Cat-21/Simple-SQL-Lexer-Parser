// timers are for future use
/*package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C //waits for timer1 to expire
	fmt.Println("Timer 1 fired")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	//here even before the timer can expire we forcefully stop the timer
	//and catch the status in the stop variable and proceed
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}*/

// tickers are used when you want to perform a task repeatedly
/*package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}*/

//panic is a runtime error that indicates that program has encountered an unexpected condition or state that it cannot handle
//w2hen called panic() immediatelt stops the normal execution of the program and starts unwinding stack abruptly stops the program
// and excuting any deferred functions along the way

/*package main

import "fmt"

func main() {
	num := 10
	denom := 0
	if denom == 0 {
		panic("bruhhh denomi is zeroo")
	}
	res := num / denom
	fmt.Println("result is ", res)

}*/

// defer is a keyword used to schedule a funnction to be executed just before the surrounding function retuens
// it is added to the deferred stack functiion calls associated with the current go routine
// used for cleanup purposes
// LIFO ORDER EXECUTION
/*package main

import "fmt"

func main() {
	defer fmt.Println("WORLD")
	fmt.Println("HELLO")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i + 1)
	}
}*/

//recover() is used to handle panics and to regain control of a panicking go routine. it is used in combination with defer
// to recover from defer gracefully and to allow a program to continue execution even after the panic occurs
// it is effective only when it is called directly from a deferred fucntion during a panic else it will retuen nil

/*package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			//panic value is captured and is assigned to r
			fmt.Println("recovered from panic", r)
		} else {
			fmt.Println("bruh simplyy")
		}
	}()
	panic("oops sommething went wrong")
}*/

