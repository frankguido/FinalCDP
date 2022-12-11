package main

import "github.com/novalagung/golpal"
import "fmt"

func main() {
	cmdString := `3 + 2`
	output, err := golpal.New().ExecuteSimple(cmdString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result", "=>", output)
}
