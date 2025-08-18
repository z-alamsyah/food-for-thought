package src

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSimpleChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Andre"
		fmt.Println("Exit GoRoutine Function")
	}()

	fmt.Println("Start Call Channel")
	result := <-channel
	fmt.Println("Result: ", result)
	fmt.Println("End Call Channel")
}

func NumberMultipliers(c chan<- int32, x int32, y int32) {
	time.Sleep(2 * time.Second)
	c <- x * y
}

func GetMultiplyResult(c <-chan int32) int32 {
	result := <-c
	return result
}

func TestChannelInParams(t *testing.T) {
	channel := make(chan int32)
	defer close(channel)

	fmt.Println("Start Calculate")
	go NumberMultipliers(channel, 12, 9)
	result := GetMultiplyResult(channel)
	fmt.Println("Result: ", result)
	fmt.Println("End Calculate")
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3) // Set Channel Capacity / Buffer is 3
	defer close(channel)

	fmt.Println("Mulai")

	channel <- "Andre"
	channel <- "Budi"
	channel <- "Wawan"
	// channel <- "Eko" // This code would make panic, caused channel buffer set is only 3 capacity

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// fmt.Println(<-channel) // This code would make panic, caused channel buffer set is only 3 capacity

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func ChannelWithRange(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- strconv.Itoa(i)
	}

	close(c) //--> For stop looping with range. The Function Close should be here, using defer would make looping forever at get value from channel with range process
}
func TestGetChannelWithRange(t *testing.T) {
	channel := make(chan string)
	// defer close(channel) //Cannot using defer if we want to use range for get the value from channel

	fmt.Println("Start")
	go ChannelWithRange(channel)
	for data := range channel {
		fmt.Println("Get Value Channel ke - ", data)
	}
	fmt.Println("End")
}

func GiveResponse(c chan<- string) {
	time.Sleep(1 * time.Second)
	c <- "Hi Bro!"
}

func TestGetChannelWithSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveResponse(channel1)
	go GiveResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Response From Channel 1 - ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Response From Channel 2 - ", data)
			counter++
		default:
			fmt.Println("Waiting for data...")
		}

		if counter == 2 {
			break
		}
	}

}
