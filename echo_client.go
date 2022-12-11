package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {

		fmt.Println("Usage: echo_client message")
		os.Exit(1)
	}

	msg := os.Args[1]

	con, err := net.Dial("udp", "debian:7")

	checkErr(err)

	defer con.Close()

	_, err = con.Write([]byte(msg))

	checkErr(err)

	reply := make([]byte, 1024)

	_, err = con.Read(reply)

	checkErr(err)

	fmt.Println("reply:", string(reply))
}

func checkErr(err error) {

	if err != nil {

		log.Fatal(err)
	}
}
