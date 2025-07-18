package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Scan(&c)
	if err != nil {
		log.Fatal(err)
	}
	if float64((a*2+b*3+c*4)/(a+b+c)) >= 3.5 {
		fmt.Println(0)
	} else {
		n := float64(1.5*float64(a)+0.5*float64(b)-0.5*float64(c)) / 1.5
		fmt.Println(math.Ceil(n))
	}

}
