package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {
    // create a handler for root
    http.HandleFunc("/", helloHandler)

    // start server
    fmt.Println("Starting Server")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
    // return hello world
    io.WriteString(w, "<h1> Hello, World! </h1>")
}
