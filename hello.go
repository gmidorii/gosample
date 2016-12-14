package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const s string = "CONST STRING"

func main() {
	fmt.Printf("hello world!")

	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

	// Type
	result := Sqrt(9.0)
	fmt.Printf("result: %v \n", result)

	// Image
	m := Image{}
	fmt.Printf(m.Bounds().String())

	// Channel
	c := make(chan int, 10)
	Send(c)
	for i := range c {
		fmt.Println(i)
	}

	ch := make(chan int)
	q := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		q <- 0
	}()
	SendSelect(ch, q)

	//const
	fmt.Println(s)

	//Closure
	f := IntFunc(10)
	fmt.Printf(strconv.Itoa(f()))

	//Interface
	rect := rect{10.0, 20.0}
	printGeo(rect)

	//File
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dirname := filepath.Join(cwd, "file")
	err = os.Mkdir(dirname, 0775)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(dirname)

	filename := filepath.Join(dirname, "tex.md")
	content := []byte("hello write\n")
	ioutil.WriteFile(filename, content, 0644)
}

/**
Interface
*/
type geometry interface {
	area() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func printGeo(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

/**
Closure
*/
func IntFunc(i int) func() int {
	return func() int {
		return i
	}
}

/**
Handler
*/
func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

/**
Type Error
*/
type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(*e))
}

func (e *ErrNegativeSqrt) String() string {
	return fmt.Sprintf("%v", *e)
}

// Sqrt
func Sqrt(x ErrNegativeSqrt) ErrNegativeSqrt {
	z := 1.0
	zNew := 1.0
	if x < 0 {
		fmt.Printf(x.String())
		return x
	}
	for {
		z = zNew
		zNew = z - ((z*z - float64(x)) / (2 * z))
		if math.Abs(zNew-z) < 1.0e-6 {
			break
		}
	}
	return ErrNegativeSqrt(zNew)
}

/**
Image
*/
type Image struct{}

func (im *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im *Image) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{0, 0},
		image.Point{200, 200},
	}
}

func (im *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x % 256), uint8(y % 256), 255, 255}
}

/**
Channel
*/
func Send(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func SendSelect(ch, q chan int) {
	for {
		select {
		case <-ch:
			fmt.Println("c")
		case <-q:
			fmt.Println("quit")
			return
		}
	}
}
