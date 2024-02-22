package game

import (
    "sync"
    "strconv"
    "math/rand/v2"
)

type state byte
type Message byte

const (
    open        state = 0
    closed      state = 1
    attempted   state = 2

    ValidOpen       Message = 0
    InvalidOpen     Message = 1
    GameWon         Message = 2
    GameOver        Message = 3
    InvalidMessage  Message = 4
)

type game struct {
    board       [][]state
    pattern     [][]bool
    mistakes    byte
    closed      byte
    tiles       byte
}

type SafeGame struct {
    mu      sync.Mutex
    wrapped game
}

var SG = SafeGame {
    wrapped: newGame(size),
}

var size        = 5
var locations   = []int{}

func (s *SafeGame) SendGame(max int, streak int) []byte {
    msg := []byte("open:")

    msg = append(msg, []byte(strconv.Itoa(max))...)
    msg = append(msg, 32)
    msg = append(msg, []byte(strconv.Itoa(streak))...)

    for i, matrix := range s.wrapped.board {
        for j, state := range matrix {
            if state == closed || state == attempted {
                if state == closed {
                    msg = append(msg, 44, 121, byte(i) + 48, 32, byte(j) + 48)
                } else {
                    msg = append(msg, 44, 110, byte(i) + 48, 32, byte(j) + 48)
                }
            }
        }
    }

    return msg 
}

func (s *SafeGame) RestartGame(won bool, max int, streak int) []byte {
    s.mu.Lock()
    s.wrapped = newGame(size)

    msg := []byte{}
    
    if won {
        msg = append(msg, []byte("won:")...)
    } else {
        msg = append(msg, []byte("new:")...)
    }

    msg = append(msg, []byte(strconv.Itoa(max))...)
    msg = append(msg, 32)
    msg = append(msg, []byte(strconv.Itoa(streak))...)

    for i, matrix := range s.wrapped.pattern {
        for j, isOpen := range matrix {
            if isOpen {
                msg = append(msg, 44, byte(i) + 48, 32, byte(j) + 48)
            }
        }
    }

    s.mu.Unlock()
    return msg
}

func (s *SafeGame) Open(message []byte) Message {
    if len(message) < 3 {
        return InvalidMessage
    }

    s.mu.Lock()

    y, x := message[0] - 48, message[2] - 48

    if !s.wrapped.pattern[y][x] {
        s.wrapped.board[y][x] = attempted
        s.wrapped.mistakes += 1

        s.mu.Unlock()
        if s.wrapped.mistakes >= 3 {
            return GameOver
        }

        return InvalidOpen
    }

    s.wrapped.board[y][x] = closed
    s.wrapped.closed += 1

    s.mu.Unlock()
    if s.wrapped.closed == s.wrapped.tiles {
        return GameWon
    } else {
        return ValidOpen
    }
}

func newGame(size int) game {
    board := make([][]state, size)
    pattern := make([][]bool, size)

    for i := range board {
        board[i] = make([]state, size)
        pattern[i] = make([]bool, size)
    }

    for i := len(locations); i < size * size; i++ {
        locations = append(locations, i)
    }

    game := game {
        board:      board,
        pattern:    pattern,
        mistakes:   0,
        closed:     0,
        tiles:      0,
    }

    game.tiles = game.generatePattern()

    return game
}

func (g *game) generatePattern() byte {
    tiles := byte(8 + rand.IntN(2))
    dropped := []int{}

    for range tiles {
        loc := rand.IntN(len(locations))
        val := drop(loc)

        y := val / size 
        x := val - size * y

        g.pattern[y][x] = true
        dropped = append(dropped, val)
    }

    for _, val := range dropped {
        locations = append(locations, val)
    }

    return tiles
}

func drop(loc int) int {
    element := locations[loc]

    locations[loc] = locations[len(locations)-1]
    locations[len(locations)-1] = 0
    locations = locations[:len(locations)-1]

    return element
}
