package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	args := os.Args // commandline arguments

	http.Handle("/", http.FileServer(http.Dir("."))) // handler for handler for server

	if len(args) == 1 { // check for port number

		go fmt.Println("Serving HTTP on 0.0.0.0 port 8000 (http://0.0.0.0:8000/) ...")
		log.Fatal(http.ListenAndServe(":8000", nil))
		os.Exit(0)

	} else {

		port, err := strconv.Atoi(args[1]) // checking if port provided is a number
		if err != nil {

			fmt.Println("usage: " + os.Args[0] + " <port>") // print usage error
			os.Exit(1)

		} else {

			// concurrent message for starting server
			go fmt.Printf("Serving HTTP on 0.0.0.0 port %d (http://0.0.0.0:%d/) ...\n", port, port)
			log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))

		}
	}
}
