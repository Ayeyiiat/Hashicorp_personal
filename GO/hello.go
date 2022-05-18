package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //Chnaging the seed value every run so as to get a different favorite number
	fmt.Println("Hello, Ayeyi.")
	fmt.Println("Welcome to Go.")
	fmt.Println("The time is now", time.Now())
	fmt.Println("My favorite random number is", rand.Intn(10))
}
