package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	usage := "usage: " + os.Args[0] + " <port>" // usage message for errors

	args := os.Args // commandline arguments

	if len(args) != 2 { // checking for only 2 arguments provided

		fmt.Println(usage) // print usage error
		os.Exit(1)

	} else {

		port, err := strconv.Atoi(args[1]) // checking if port provided is a number
		if err != nil {

			fmt.Println(usage) // print usage error
			os.Exit(1)

		} else {

			http.Handle("/", http.FileServer(http.Dir("."))) // handler for handler for server
			go fmt.Printf("Serving HTTP on 0.0.0.0 port %d (http://0.0.0.0:%d/) ...\n", port, port) // concurrent message for starting server
			log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))

		}
	}
}
