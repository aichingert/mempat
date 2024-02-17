package game

type Game struct {
    fields [4][4]bool
}

func New() Game {
    game := Game {
        fields: [4][4]bool {
            {false,false,false,false},
            {false,false,false,false},
            {false,false,false,false},
            {false,false,false,false},
        },
    }

    return game
}
