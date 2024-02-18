package game

type Game struct {
    board [][]bool
    pattern [][]bool
}

var G = New()

func New() Game {
    game := Game {
        board: [][]bool {{3: false}, {3: false}, {3: false}, {3: false}},
        pattern: generatePattern(),
    }

    return game
}

func (g *Game) SendGame() []byte {
    board := []byte("open:")

    for i, c := range g.board {
        for j, _ := range c {
            if c[j] {
                if len(board) > 5 {
                    board = append(board, 44)
                }
                board = append(board, byte(i) + 48, 32, byte(j) + 48)
            }
        }
    }

    return board
}

func (g *Game) Open(message []byte) bool {
    if len(message) < 3 {
        return false;
    }

    y, x := message[0] - 48, message[2] - 48

    if !g.pattern[y][x] {
        return false;
    }

    g.board[y][x] = true
    return true;
}

func generatePattern() [][]bool {
    return [][]bool{
        {false,false,true,false},
        {true,true,false,false},
        {false,false,true,false},
        {false,false,false,false},
    }
}
