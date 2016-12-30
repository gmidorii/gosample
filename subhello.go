package main

import (
	"time"
	"fmt"
)

func main() {
	slice := make([]int, 100000000)
	for i, _ := range slice {
		slice[i] = i
	}

	// normal
	var sum int
	start := time.Now()
	for _, v := range slice {
		sum += v
	}
	fmt.Println(sum)
	end := time.Now()
	fmt.Printf("%f sec \n", end.Sub(start).Seconds())

	// channel
	sum = 0
	startT := time.Now()
	c1 := make(chan int)
	c2 := make(chan int)
	c1slice := slice[:len(slice)/2]
	c2slice := slice[len(slice)/2:]
	Sum(c1, c1slice)
	Sum(c2, c2slice)
	sum += <- c1
	sum += <- c2
	fmt.Println(sum)
	endT := time.Now()
	fmt.Printf("%f sec \n", endT.Sub(startT).Seconds())
}

func Sum(c chan int, s []int) {
	go func() {
		sum := 0
		for _, v := range s {
			sum += v
		}
		c <- sum
	}()
}
