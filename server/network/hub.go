package network

import (
    "log"
    "sync"
    "mempat/game"
)

type SafeBuffer struct {
    mu          sync.Mutex
    msg         []byte
}

var SB = SafeBuffer {
    msg:        []byte{},
}

type Hub struct {
    max         int
    streak      int
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
            SB.mu.Lock()
            h.generateMessage(position)

            for client := range h.clients {
                select {
                case client.send <- SB.msg:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }

            SB.mu.Unlock()
        }
    }
}

func (h *Hub) generateMessage(position []byte) {
    SB.msg = []byte{}

    switch status := game.SG.Open(position); status {
    case game.ValidOpen:
        //                  v,   a,   l,  :
        SB.msg = append(SB.msg, 118,  97, 108, 58)
        SB.msg = append(SB.msg, position...)
    case game.InvalidOpen:
        //                  i,   n,   v,  :
        SB.msg = append(SB.msg, 105, 110, 118, 58)
        SB.msg = append(SB.msg, position...)
    case game.GameWon:
        h.streak += 1
        h.max = max(h.streak, h.max)

        SB.msg = append(SB.msg, game.SG.RestartGame(true, h.max, h.streak)...)
    case game.GameOver:
        h.streak = 0
        SB.msg = append(SB.msg, game.SG.RestartGame(false, h.max, h.streak)...)
    default:
        log.Println("ERROR: ", status)
    }
}
