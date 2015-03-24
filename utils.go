package levelgen

import (
	"fmt"
	"os"
	"github.com/ajstarks/svgo"
)

func MakePatternStyle(w int, h int, pct int, canvas *svg.SVG) string {
	pw, ph := (w*pct)/100, (h*pct)/100
	name := "hatch"
	canvas.Def()
	canvas.Pattern(name, 0, 0, pw, ph, "user")
	canvas.Gstyle("fill:none;stroke-width:1")
	canvas.Path(fmt.Sprintf("M0,0 l%d,%d", pw, ph), "stroke:green")
	canvas.Path(fmt.Sprintf("M%d,0 l-%d,%d", pw, pw, ph), "stroke:blue")
	canvas.Gend()
	canvas.PatternEnd()
	canvas.DefEnd()

	return fmt.Sprintf("fill:url(#%s)", name)
}

// Return int if greater than v
// else return v
//
func bi(v int, n int) int {
	if v < n {
		return n
	}
	return v
}

func MakeFileCanvas(w int, h int, filename string) *svg.SVG {
	canvas := svg.New(GetWriter(filename))
	canvas.Start(w, h)
	return canvas
}

// Create a file writer
//
func GetWriter(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
   		fmt.Fprintf(os.Stderr, "%v\n", err)
   		return os.Stdout
   	}
   	return f
}