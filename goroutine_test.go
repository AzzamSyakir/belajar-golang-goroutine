package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("hello world")
}
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("ups")

}
func DisplayNumber(number int) {
	fmt.Println("display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
