package main

import (
    "fmt"
    "log"
    "time"
    "net/http"
    "multiplayer/network"
)

func main() {
    fmt.Println("hello, world!")

    network.SetupRoutes()

    server := &http.Server {
        Addr:              ":8080",
        ReadHeaderTimeout: 3 * time.Second,
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
