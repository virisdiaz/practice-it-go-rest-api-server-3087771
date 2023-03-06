package main

import (
	"fmt"
	"log"

	"net/http"

	"viridiana.com/viri/practice"
)

func main() {
	http.HandleFunc("/", practice.HelloWorld)
	fmt.Println("Server started listening on localhost:9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}
