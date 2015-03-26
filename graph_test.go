package levelgen

import (
	"github.com/highway900/goNetwork"
	"testing"
	"math/rand"
	"time"
	"fmt"
	"github.com/ajstarks/svgo"
)

func TestGraphDraw(t *testing.T) {

	fmt.Println("Draw graph Test")

	rand.Seed(time.Now().Unix())
	
	canvas := svg.New(GetWriter("test_graph_draw.svg"))

	w := 500
	h := 300

	canvas.Start(w, h)

	graph := new(goNetwork.Graph)

	// create a node
	node1 := graph.AddNode(MakeCircle(rand.Intn(w), rand.Intn(h), bi(3, rand.Intn(14))))
	node2 := graph.AddNode(MakeCircle(rand.Intn(w), rand.Intn(h), bi(3, rand.Intn(14))))
	node3 := graph.AddNode(MakeCircle(rand.Intn(w), rand.Intn(h), bi(3, rand.Intn(14))))

	graph.Connect(node1, node2, 1.0, "k")
	graph.Connect(node1, node3, 1.0, "k")

	for _, edge := range graph.Edges {
		line := MakeLine(
			edge.P0().Data.(Shape).Centre(),
			edge.P1.Data.(Shape).Centre(),
		)
		line.Svg(canvas)
	}

	for _, node := range graph.Nodes {
		node.Data.(Shape).Svg(canvas)
	}

	canvas.End()

}