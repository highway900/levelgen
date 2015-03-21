package levelgen

import (
    "testing"
    "fmt"
    "os"
    "math/rand"
    "time"
	"github.com/ajstarks/svgo"
)

func ShapeTester(s Shape, canvas *svg.SVG) {
	s.Svg(canvas)
	s.Centre()	
}

func GetWriter(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
   		fmt.Fprintf(os.Stderr, "%v\n", err)
   		return os.Stdout
   	}
   	return f
}

func TestCircle2(t *testing.T) {
	fmt.Println("Testing Circle Pattern")

	canvas := svg.New(GetWriter("test_circle_pattern.svg"))
	
	w, h := 500, 500
	pct := 2
	pw, ph := (w*pct)/100, (h*pct)/100
	canvas.Start(w, h)
	
	// define the pattern
	canvas.Def()
	canvas.Pattern("hatch", 0, 0, pw, ph, "user")
	canvas.Gstyle("fill:none;stroke-width:1")
	canvas.Path(fmt.Sprintf("M0,0 l%d,%d", pw, ph), "stroke:red")
	canvas.Path(fmt.Sprintf("M%d,0 l-%d,%d", pw, pw, ph), "stroke:blue")
	canvas.Gend()
	canvas.PatternEnd()
	canvas.DefEnd()

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