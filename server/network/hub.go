package network

import (
    "log"
    "sync"
    "mempat/game"
)

type Hub struct {
    max         int
    streak      int
    mu          sync.Mutex
    clients     map[*Client]bool
    broadcast   chan []byte

    register    chan *Client
    unregister  chan *Client
}

func NewHub() *Hub {
    return &Hub {
        max:        0,
        streak:     0,
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
            msg := h.generateMessage(position)

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

func (h *Hub) generateMessage(position []byte) []byte {
    h.mu.Lock()
    defer h.mu.Unlock()
    msg := []byte{}

    switch status := game.SG.Open(position); status {
    case game.ValidOpen:
        //                  v,   a,   l,  :
        msg = append(msg, 118,  97, 108, 58)
        msg = append(msg, position...)
    case game.InvalidOpen:
        //                  i,   n,   v,  :
        msg = append(msg, 105, 110, 118, 58)
        msg = append(msg, position...)
    case game.GameWon:
        h.streak += 1
        h.max = max(h.streak, h.max)

        msg = append(msg, game.SG.RestartGame(true, h.max, h.streak)...)
    case game.GameOver:
        h.streak = 0
        msg = append(msg, game.SG.RestartGame(false, h.max, h.streak)...)
    default:
        log.Println("ERROR: ", status)
    }

    return msg
}
