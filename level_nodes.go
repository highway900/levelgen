package levelgen

import (
	"fmt"
	"math/rand"
	"github.com/highway900/goNetwork"
	"github.com/ajstarks/svgo"
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
	pt := &Point{
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

func (l *Level) Draw(filename string) {

	// TODO: Fix below
	canvas := MakeFileCanvas(l.Width, l.Height, filename)


	style := &SVGStyle{25, "blue", ""}

	for _, edge := range l.graph.Edges {
		line := MakeLine(
			edge.P0().Data.(*LevelNode).GetPoint(),
			edge.P1.Data.(*LevelNode).GetPoint(),
		)
		line.SetStyle(style.MakeStyle())
		line.Svg(canvas)
	}

	style = &SVGStyle{5, "black", "green"}
	for _, node := range(l.graph.Nodes) {
		shape := node.Data.(*LevelNode).GetShape()
		shape.SetStyle(style.MakeStyle())
		node.Data.(*LevelNode).Draw(canvas)
	}

	canvas.End()

}

type ILevelNode interface {
	Action()
}

type LevelNode struct {
	pt *Point
	value int
	shape Shape
}

func (ln *LevelNode) Draw(canvas *svg.SVG) {
	ln.shape.Svg(canvas)
}

func (ln *LevelNode) GetPoint() *Point {
	return ln.pt
}

func (ln *LevelNode) GetShape() Shape {
	return ln.shape
}

func (ln *LevelNode) Action() {
	// Do something exciting here
	fmt.Println("ACTION", ln.value)
}

func MakeLevelNode(pt *Point, value int) *LevelNode {
	return &LevelNode{
		pt,
		value,
		MakeCirclePt(
			pt,
			bi(50, rand.Intn(100)),
		),
	}
}