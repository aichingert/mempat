package game

import "math/rand/v2"

type state byte
type Message byte

const (
    open        state = 0
    closed      state = 1
    attempted   state = 2

    ValidOpen       Message = 0
    InvalidOpen     Message = 1
    GameOver        Message = 2
    InvalidMessage  Message = 3
)

type Game struct {
    board [][]state
    pattern [][]bool
    mistakes byte
}

var G = NewGame(size)

var size  = 5
var locations = []int{}

func (g *Game) SendGame() []byte {
    msg := []byte("open:")

    for i, matrix := range g.board {
        for j, state := range matrix {
            if state == closed || state == attempted {
                if len(msg) > 5 {
                    msg = append(msg, 44)
                }

                if state == closed {
                    msg = append(msg, 121, byte(i) + 48, 32, byte(j) + 48)
                } else {
                    msg = append(msg, 110, byte(i) + 48, 32, byte(j) + 48)
                }
            }
        }
    }

    return msg 
}

func (g *Game) RestartGame() []byte {
    G = NewGame(size)

    msg := []byte("new:")

    for i, matrix := range g.pattern {
        for j, isOpen := range matrix {
            if isOpen {
                if len(msg) > 5 {
                    msg = append(msg, 44)
                }
                msg = append(msg, byte(i) + 48, 32, byte(j) + 48)
            }
        }
    }

    return msg
}

func (g *Game) Open(message []byte) Message {
    if len(message) < 3 {
        return InvalidMessage
    }

    y, x := message[0] - 48, message[2] - 48

    if !g.pattern[y][x] {
        g.board[y][x] = attempted
        g.mistakes += 1

        if g.mistakes >= 3 {
            return GameOver
        }

        return InvalidOpen
    }

    g.board[y][x] = closed
    return ValidOpen 
}

func NewGame(size int) Game {
    board := make([][]state, size)
    pattern := make([][]bool, size)

    for i := range board {
        board[i] = make([]state, size)
        pattern[i] = make([]bool, size)
    }

    for i := len(locations); i < size * size; i++ {
        locations = append(locations, i)
    }

    game := Game {
        board: board,
        pattern: pattern,
        mistakes: 0,
    }

    game.generatePattern()

    return game
}

func (g *Game) generatePattern() {
    dropped := []int{}

    for range 8 + rand.IntN(2) {
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
}

func drop(loc int) int {
    element := locations[loc]

    locations[loc] = locations[len(locations)-1]
    locations[len(locations)-1] = 0
    locations = locations[:len(locations)-1]

    return element
}
