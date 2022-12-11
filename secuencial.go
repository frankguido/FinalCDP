package main

import (
	"fmt"
	"log"
	"time"
)

func mostrar0() {
	for f := 1; f <= 1000; f++ {
		fmt.Print("0-")
	}
}

func mostrar1() {
	for f := 1; f <= 1000; f++ {
		fmt.Print("1-")
	}
}

func main() {
	start := time.Now()
	mostrar0()
	mostrar1()
	elapsed := time.Since(start)
	log.Printf("tiempo es %s :", elapsed)
}
