package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func useTime() {
	fmt.Println("Time Now:", time.Now())

	var timeFormat = time.Date(2023, 06, 24, 12, 0, 0, 0, time.Local)
	fmt.Println("Time Format:", timeFormat)
}

func useTimeNow() {
	var now = time.Now()

	fmt.Println("Year:", now.Year(), "Month:", now.Month(), "Day:", now.Day())
}

func timeParse() {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2023-10-02 08:10:05"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t ->", date.String())

	layoutFormat = "2006-01-02 MST"
	value = "2023-01-02 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t ->", date.String())
}

func predefinedLayout() {
	var date, _ = time.Parse(time.RFC850, "Sunday, 01-Jan-08 12:02:00 WIB")
	fmt.Println(date)
}

// Timer,Ticker,& Scheduler

func useBasicScheduler() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello!")
		time.Sleep(time.Second)
	}
}

func useNewTimer() {
	newTime := time.NewTimer(time.Second * 5)

	fmt.Println("Start")
	<-newTime.C
	fmt.Println("End")
}

func useAfterFunc() {
	var ch = make(chan bool)

	time.AfterFunc(time.Second*5, func() {
		fmt.Println("Timeout!")
		ch <- true
	})

	fmt.Println("Start")
	<-ch
	fmt.Println("End")
}

func useAfter() {
	fmt.Println("Start")
	<-time.After(time.Second * 5)
	fmt.Println("End")
}

func useTicker() {
	var done = make(chan bool)
	var ticker = time.NewTicker(time.Second)

	go func() {
		time.Sleep(time.Second * 10)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			fmt.Println("Stop")
			return
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}
}

func timing(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Second*time.Duration(timeout), func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\nNot answered after", timeout, ". Quiz Ended!")
	os.Exit(0)
}

func timerNGoroutine() {
	var ch = make(chan bool)
	var timeout int = 10

	go timing(timeout, ch)
	go watcher(timeout, ch)

	for {
		var input string
		var input1, input2, val int
		input1 = rand.Int()%100 + 1
		input2 = rand.Int()%100 + 1
		fmt.Print("Whats is ", input1, " + ", input2, ": ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Print("Invalid value, please input number!\n\n")
			continue
		}

		val, err = strconv.Atoi(input)
		if err != nil {
			fmt.Print("Invalid value, please input number!\n\n")
			continue
		}

		if val == input1+input2 {
			fmt.Println("Answer is correct!")
			break
		} else {
			fmt.Print("Answer is wrong! Try Again\n\n")
			continue
		}
	}

}

func main() {
	//useTime()
	//useTimeNow()
	//timeParse()
	//predefinedLayout()
	//useBasicScheduler()
	//useNewTimer()
	//useAfterFunc()
	//useAfter()
	//useTicker()
	timerNGoroutine()
}
