package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for i := 1; i <= 1000; i++ {
		_, err2 := f.WriteString(strconv.Itoa(i) + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
	}

	fmt.Println("done")
}
