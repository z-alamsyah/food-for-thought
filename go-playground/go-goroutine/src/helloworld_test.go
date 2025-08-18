package src

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello Folks!")
}

func TestHelloWorld(t *testing.T) {
	fmt.Println("Start")
	go HelloWorld()
	fmt.Println("End")

	time.Sleep(1 * time.Second)
}
