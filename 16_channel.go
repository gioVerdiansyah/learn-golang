package main

import (
	"first_golang/my_modules"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func useChannel() {
	message := make(chan string)

	var sayHello = func(who string) {
		var data = fmt.Sprintf("Hello, %s!", who)
		message <- data
	}

	go sayHello("Sofyan Gio")
	go sayHello("Alex Verdi")
	go sayHello("Adi Niam")

	var msg1 = <-message
	fmt.Println("1.", msg1)
	var msg2 = <-message
	fmt.Println("2.", msg2)
	var msg3 = <-message
	fmt.Println("3.", msg3)
}

func channelAsParam() {
	var messages = make(chan string)

	var printMessage = func(data chan string) {
		fmt.Println(<-data)
	}

	var dataMsg = []string{"Verdi", "Niam", "Adi"}

	for _, message := range dataMsg {
		go func(msg string) {
			data := fmt.Sprintf("Hello %s!", msg)
			messages <- data
		}(message)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

func bufferedChannel() {
	messages := make(chan int, 4)

	go func() {
		for {
			i := <-messages
			fmt.Printf("Receive %d\n", i)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("Sending %d\n", i)
		messages <- i
	}
}

func getAvg(val []int, ch chan int) {
	ch <- my_modules.Avg(val)
}

func getMax(val []int, ch chan int) {
	ch <- my_modules.Max(val)
}

func getMin(val []int, ch chan int) {
	ch <- my_modules.Min(val)
}

func selectChannel() {
	nums := []int{2, 2, 3, 3, 3, 3, 4, 5, 5}
	fmt.Printf("Value: %v\n", nums)

	var ch1 = make(chan int)
	go getAvg(nums, ch1)

	var ch2 = make(chan int)
	go getMax(nums, ch2)

	var ch3 = make(chan int)
	go getMin(nums, ch3)

	for i := 0; i < 3; i++ {
		select {
		case avgVal := <-ch1:
			fmt.Printf("Avg: %v\n", avgVal)
		case maxVal := <-ch2:
			fmt.Printf("Max: %v\n", maxVal)
		case minVal := <-ch3:
			fmt.Printf("Min: %v\n", minVal)
		}
	}
}

func sendMessage(ch chan<- string) {
	for i := 1; i <= 10; i++ {
		dataMsg := fmt.Sprintf("Hello index %d", i)
		ch <- dataMsg
	}

	close(ch)
}

func receiveMessage(ch <-chan string) {
	for val := range ch {
		fmt.Println(val)
	}
}

func forRangeClose() {
	messages := make(chan string)
	go sendMessage(messages)
	receiveMessage(messages)
}

func sendTimeoutData(ch chan<- int) {
	for i := 1; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%15+1) * time.Second)
	}
}

func receiveTimeoutData(ch <-chan int) {
outerLoop:
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-time.After(5 * time.Second):
			fmt.Print("Not Receive data after 5 seconds. Break loop...")
			break outerLoop
		}
	}
}

func channelTimeout() {
	var messages = make(chan int)

	go sendTimeoutData(messages)
	receiveTimeoutData(messages)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//useChannel()
	//channelAsParam()
	//bufferedChannel()
	//selectChannel()
	//forRangeClose()
	channelTimeout()

	var input string
	fmt.Scanln(&input)
}

// Channel Directions
// ch chan string 	| Sending and Receive
// ch chan<- string | Sending
// ch <-chan string | Receive
