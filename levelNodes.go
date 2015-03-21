package levelgen

import (
	"fmt"
	"math/rand"
	"github.com/highway900/goNetwork"
)

// Effectively a wrapper on a goNetwork graph
// 
type Level struct {
	seed int64
	Width int
	Height int
	graph *goNetwork.Graph
}

func MakeLevel(w int, h int, seed int64) *Level {
	return &Level{
		seed,
		w,
		h,
		new(goNetwork.Graph),
	}
}

func (l *Level) createRandomNode() {
	pt := Point{
		rand.Intn(l.Width),
		rand.Intn(l.Height),
	}
	l.graph.AddNode(MakeLevelNode(pt, 1))
}

// nn: Number of nodes to generate
// seed: Random seed value
// Example -> level.GenSimpleRandLvl(10, time.Now().Unix())
//
func (l *Level) GenSimpleRandomLVL(nn int, seed int64) {
	l.graph = new(goNetwork.Graph)
	l.seed = seed
	
	rand.Seed(seed)

	// Create some random nodes
	for i := 0; i<nn; i++ {
		l.createRandomNode()		
	}

	// Connect the nodes
	for i, node := range(l.graph.Nodes) {
		if i != 0 {
			prevNode := l.graph.Nodes[i-1]
			l.graph.Connect(prevNode, node, rand.Float64(), fmt.Sprintf("%d", i-1))
		}
	}
}

type ILevelNode interface {
	Action()
}

type LevelNode struct {
	pt Point
	value int
}

func (ln *LevelNode) Action() {
	// Do something exciting here
}

func MakeLevelNode(pt Point, value int) *LevelNode {
	return &LevelNode{
		pt,
		value,
	}
}