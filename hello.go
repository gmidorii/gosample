package main

import (
	"fmt"
	"net/http"
	"math"
)


type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(*e));
}

func (e *ErrNegativeSqrt) String() string {
	return fmt.Sprintf("%v", *e)
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func main() {
	result := Sqrt(9.0)
	fmt.Printf("result: %v \n", result)
	fmt.Printf("hello world!")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func Sqrt(x ErrNegativeSqrt) ErrNegativeSqrt {
	z := 1.0
	zNew := 1.0
	if x < 0 {
		fmt.Printf(x.String())
		return x
	}
	for {
		z = zNew
		zNew = z - ((z * z - float64(x)) / (2 * z))
		if math.Abs(zNew - z) < 1.0e-6 {
			break
		}
	}
	return ErrNegativeSqrt(zNew)
}

