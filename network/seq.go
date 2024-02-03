package network

type Coord struct {
    x int
    y int
}

type Game struct {
    seq []Coord
    current int
}

// TODO: Implement the missing logic for this
//
// 1: Being able to progress in the sequence by pressing a button
// 2: Create the next element in the sequence if you reached the last one

func (g *Game) Check() bool {
    return len(g.seq) == g.current
}
