package network

import "multiplayer/game"

type Hub struct {
    game        chan game.Game
    clients     map[*Client]bool
    broadcast   chan []byte

    register    chan *Client
    unregister  chan *Client
}

func NewHub() *Hub {
    return &Hub {
        game:       make(chan game.Game),
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
        case message := <-h.broadcast:
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
        }
    }
}
