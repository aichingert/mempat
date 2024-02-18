package main

import (
    "log"
    "time"
    "net/http"
    "mempat/network"
)

func SetupRoutes() {
    hub := network.NewHub() ; go hub.Run()

    // hosting the client site
    http.Handle("/", http.FileServer(http.Dir("../client")))

    // managing websocket connections
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        network.ServeWs(hub, w, r)
    })
}

func main() {
    SetupRoutes()

    addr := ":8080"
    server := &http.Server {
        Addr:              addr,
        ReadHeaderTimeout: 3 * time.Second,
    }

    log.Printf("Running on localhost%s\n", addr)

    err := server.ListenAndServe()
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
