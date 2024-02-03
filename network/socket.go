package network 

import (
    "net/http"
)

var game = Game {
    seq: []Coord{},
    current: 0,
}

func SetupRoutes() {
    hub := newHub()
    fs := http.FileServer(http.Dir("./client"))

    http.Handle("/", fs)
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(hub, w, r)
    })
}
