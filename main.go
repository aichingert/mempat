package main

import (
    "log"
    "time"
    "net/http"
    "multiplayer/network"
)

var game = network.Game {
    Seq: []network.Coord{},
    Current: 0,
}

func SetupRoutes() {
    hub := network.NewHub()
    go hub.Run()
    fs := http.FileServer(http.Dir("./client"))

    http.Handle("/", fs)
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        network.ServeWs(hub, w, r)
    })
}

func main() {
    SetupRoutes()

    server := &http.Server {
        Addr:              ":8080",
        ReadHeaderTimeout: 3 * time.Second,
    }

    log.Println("Running on 8080")
    err := server.ListenAndServe()
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
