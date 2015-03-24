package levelgen

import (
	"fmt"
)

type SVGStyle struct {
	Stroke float32
	StrokeColour string

	Fill string
}

func (s *SVGStyle) MakeStyle() string {
	style := ""
	if s.StrokeColour != "" {
		style = style + fmt.Sprintf(
			"stroke: %s;", s.StrokeColour)
	}
	if s.Stroke != 0.0 {
		style = style + fmt.Sprintf(
			"stroke-width: %f;", s.Stroke)
	}
	if s.Fill != "" {
		style = style + fmt.Sprintf(
			"fill: %s;", s.Fill)
	}
	return style
}