package main

import (
    "fmt"
    "log"
    "net/http"
    "multiplayer/network"
)

func main() {
    fmt.Println("hello, world!")
    network.SetupRoutes()
    log.Fatal(http.ListenAndServe(":8080", nil))
}
