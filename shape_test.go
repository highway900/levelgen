package levelgen

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
	"github.com/ajstarks/svgo"
)

func ShapeTester(s Shape, canvas *svg.SVG) {
	s.Svg(canvas)
	s.Centre()	
}

func TestCirclePatternFill(t *testing.T) {
	fmt.Println("Testing Circle Pattern")

	canvas := svg.New(GetWriter("test_circle_pattern.svg"))
	
	w, h := 500, 500
	canvas.Start(w, h)
	
	MakePatternStyle(2, w, h, canvas)

	canvas.Gstyle("stroke:black; stroke-width:2")
	nstars := 20
	for i := 0; i < nstars; i++ {
		circle := MakeCircle(rand.Intn(w), rand.Intn(h), 15)
		circle.SetStyle("fill:url(#hatch)")
		ShapeTester(circle, canvas)
	}

	canvas.Gend()
	canvas.End()
}

func TestCircle(t *testing.T) {

	fmt.Println("Testing Circle")
	
	rand.Seed(time.Now().Unix())
	
	canvas := svg.New(GetWriter("test_circle.svg"))

	width := 500
	height := 300

	canvas.Start(width, height)

	nstars := 20
	for i := 0; i < nstars; i++ {
		c := MakeCircle(rand.Intn(width), rand.Intn(height), 15)
		ShapeTester(c, canvas)
	}

	canvas.End()

}