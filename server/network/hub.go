package network

import (
    "log"
    //"strings"
    "mempat/game"
)

type Hub struct {
    clients     map[*Client]bool
    broadcast   chan []byte

    register    chan *Client
    unregister  chan *Client
}

func NewHub() *Hub {
    return &Hub {
        clients:    make(map[*Client]bool),
        broadcast:  make(chan []byte),

        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
        case client := <-h.unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
        case position := <-h.broadcast:
            isValid := game.G.Open(position)

            msg := []byte{}

            // prefixes for valid and invalid
            // used for the colors on the client
            if isValid {
                //                  v,   a,   l,  :
                msg = append(msg, 118,  97, 108, 58)
            } else {
                //                  i,   n,   v,  :
                msg = append(msg, 105, 110, 118, 58)
            }

            msg = append(msg, position...)

            log.Println(msg)

            for client := range h.clients {
                select {
                case client.send <- msg:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
        }
    }
}
