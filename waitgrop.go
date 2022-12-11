package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func mostrar0() {
	for f := 1; f <= 1000; f++ {
		fmt.Print("0-")
	}
	wg.Done()
}

func mostrar1() {
	for f := 1; f <= 1000; f++ {
		fmt.Print("1-")
	}
	wg.Done()
}

func main() {
	start := time.Now()
	wg.Add(2)
	go mostrar0()
	go mostrar1()
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("tiempo es %s :", elapsed)
}
