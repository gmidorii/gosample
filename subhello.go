package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Test struct {
	num int
	str string
}

type Counter struct {
	num int
}

var addr = flag.String("addr", ":1718", "http service address")

func main() {
	/*
		1. parallel exec
	*/
	slice := make([]int, 100000000)
	for i, _ := range slice {
		slice[i] = i
	}
	// normal
	normSum(slice)
	// channel
	chanSum(slice)

	/*
		2. map
	*/
	var value string
	var ok bool
	tMap := map[string]string{"1": "one", "2": "two"}
	// false
	value, ok = tMap["100"]
	existPrint(ok, value)
	// true
	value, ok = tMap["2"]
	existPrint(ok, value)
	// %v is almost all format
	fmt.Printf("%v\n", tMap)

	/*
		3. String()
	*/
	t := Test{num: 1, str: "TEST"}
	fmt.Printf("%v", t)

	/*
		4. array append
	*/
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...) // "..." is necessary
	fmt.Println(x)

	/*
		5. comma, ok idiom
	*/
	var v interface{} = 0
	if str, ok := v.(string); ok {
		fmt.Println(str)
	} else {
		fmt.Println("value is not string")
	}

	/*
		6. net/http interface
	*/
	ctr := new(Counter)
	http.Handle("/counter", ctr)
	http.Handle("/arg", http.HandlerFunc(ArgServer))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("http error")
	}
}

// String return format string
// caution: Type *Test is wrong -> Type Test is correct
func (t Test) String() string {
	return fmt.Sprintf("num: %d\nstring: %s\n", t.num, t.str)
}

func (c *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c.num++
	fmt.Fprintf(w, "count: %d", c.num)
}

// ArgServer is HandlerFunction
// usage http.HandlerFunc(ArgServer)
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args)
}

func existPrint(ok bool, value string) {
	if ok {
		fmt.Printf("%s\n", value)
	} else {
		fmt.Println("no value")
	}
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

func chanSum(slice []int) {
	var sum int
	startT := time.Now()
	c1 := make(chan int)
	c2 := make(chan int)
	c1slice := slice[:len(slice)/2]
	c2slice := slice[len(slice)/2:]
	Sum(c1, c1slice)
	Sum(c2, c2slice)
	sum += <-c1
	sum += <-c2
	fmt.Println(sum)
	endT := time.Now()
	fmt.Printf("%f sec \n", endT.Sub(startT).Seconds())

}

func normSum(slice []int) {
	var sum int
	start := time.Now()
	for _, v := range slice {
		sum += v
	}
	fmt.Println(sum)
	end := time.Now()
	fmt.Printf("%f sec \n", end.Sub(start).Seconds())

}
