package levelgen

import (
	"math"
	"github.com/ajstarks/svgo"
)

type Shape interface {
	Svg(canvas *svg.SVG)
	Centre() *Point
	SetStyle(style string)
}

type Point struct {
	X int
	Y int
}

type Circle struct {
	centre *Point
	radius int
	style string
}

type Line struct {
	P1 *Point
	P2 *Point
	length float64
	centre *Point
	style string
}

func (p *Point) Minus(otherPoint *Point) *Point {
	return &Point {
		p.X - otherPoint.X,
		p.Y - otherPoint.Y,
	}
}

func (p *Point) Dot(p1 *Point) float64 {
	return float64(p.X * p1.X + p.Y * p1.Y)
}

func (p *Point) Dist(p1 *Point) float64 {
	x := p.X - p1.X
	y := p.Y - p1.Y
	return math.Sqrt(float64(x * x + y * y))
}

func MakeLine(p1 *Point, p2 *Point) *Line {
	centre := &Point{p1.X-p2.X, p1.Y-p2.Y}
	length := 1.0 // Lazy and dumb I am
	return &Line {
		p1,
		p2,
		length,
		centre,
		"fill:none;stroke:black",
	}
}

func (l *Line) Svg(canvas *svg.SVG) {
	canvas.Line(l.P1.X, l.P1.Y, l.P2.X, l.P2.Y, l.style)
}

func (l *Line) Centre() *Point {
	return l.centre
}

// This might not work as the Circle object itself is not updated
func (l *Line) SetStyle(style string) {
	l.style = style
}

func MakeCircle(x int, y int, radius int) *Circle {
	centre := &Point{x, y}
	return &Circle{
		centre,
		radius,
		"fill:none;stroke:black",
	}
}

func MakeCirclePt(pt *Point, radius int) *Circle {
	return MakeCircle(pt.X, pt.Y, radius)
}

func (c *Circle) GetRadius() int {
	return c.radius
}

func (c *Circle) Svg(canvas *svg.SVG) {
	canvas.Circle(c.centre.X, c.centre.Y, c.radius, c.style)
}

func (c *Circle) Centre() *Point {
	return c.centre
}

// This might not work as the Circle object itself is not updated
func (c *Circle) SetStyle(style string) {
	c.style = style
}

func (c *Circle) Intersects(otherCircle Circle) bool {
	// Do intersection test
	centreDiff := c.Centre().Minus(otherCircle.Centre())
	radiusSum := c.radius - otherCircle.GetRadius()
	return centreDiff.Dot(centreDiff) <= float64(radiusSum * radiusSum)
}
