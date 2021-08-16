package main

import (
	"fmt"
	"github.com/tatyanayagolnikov/new_go/helpers"
)

const x = 10

func CalcVal(c chan int) {
	rand := helpers.RandNum(x)
	c <- rand // SEND from CHANNEL 
	
}

func main() {
	c2 := make(chan string)
	c1 := make(chan int) // Make CHANNEL

	defer close(c1) // whatever comes after keyword DEFER, execute that as soon as the current function is done 
	go CalcVal(c1) 

	go helpers.Hop(c2)
	helpers.Drop(c2)


// RECEIVE
num := <-c1
fmt.Println(num)


var p helpers.Person
fmt.Println(p)
fmt.Println("Exit -Tanya Programmer")

} 