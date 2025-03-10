package main

import (
	"fmt"
	"time"
)

func looping(s string) {
	for i := range 3 {
		fmt.Printf("hello: %s ,index: %v \n", s, i)
	}
}

func main() {
	// fmt.Println("Hello Wolrd")
	go looping("Tuschy")
	looping("Nriny")

	time.Sleep(1 * time.Second)

}
