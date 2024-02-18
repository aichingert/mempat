package game

type Game struct {
    board [][]bool
    pattern [][]bool
    mistakes byte
}

var G = NewGame()

func (g *Game) SendGame() []byte {
    return generateMessageWithOpenFieldsAndPrefix([]byte("open:"), g.board)
}

func (g *Game) RestartGame() []byte {
    G = NewGame()
    return generateMessageWithOpenFieldsAndPrefix([]byte("new:"), g.pattern)
}

// TODO: maybe use an enum for return values
func (g *Game) Open(message []byte) byte {
    if len(message) < 3 {
        return 3 // invalid message
    }

    y, x := message[0] - 48, message[2] - 48

    if !g.pattern[y][x] {
        g.mistakes += 1

        if g.mistakes >= 2 {
            return 2 // game over
        }

        return 1 // invalid open
    }

    g.board[y][x] = true
    return 0 // valid open
}

func NewGame() Game {
    game := Game {
        board: [][]bool {{3: false}, {3: false}, {3: false}, {3: false}},
        pattern: generatePattern(),
        mistakes: 0,
    }

    return game
}

func generatePattern() [][]bool {
    return [][]bool{
        {false,false,true,false},
        {true,true,false,false},
        {false,false,true,false},
        {false,false,false,false},
    }
}

func generateMessageWithOpenFieldsAndPrefix(prefix []byte, matrix [][]bool) []byte {
    msg := prefix

    for i, c := range matrix {
        for j, _ := range c {
            if c[j] {
                if len(msg) > 5 {
                    msg = append(msg, 44)
                }
                msg = append(msg, byte(i) + 48, 32, byte(j) + 48)
            }
        }
    }

    return msg
}
