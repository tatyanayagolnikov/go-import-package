package helpers

import (
	"math/rand"
	"time"
	"fmt"
)

//import "fmt"

type Person struct {
	Name string
	Sport string 
}

func RandNum(n int) int {
	rand.Seed(time.Now().UnixNano()) // generates random number
	value := rand.Intn(n)
	return value
}

// SEND CHANNEL 
func Hop(c chan string) {
	c <- "pizza"
}

// RECEIVE CHANNEL
func Drop(c chan string) {
	fmt.Println(<-c)
}



