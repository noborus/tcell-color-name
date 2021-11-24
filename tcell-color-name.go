package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/gdamore/tcell/v2"
	"github.com/jwalton/gchalk"
)

type ColorName struct {
	Name  string
	Color tcell.Color
}
type Colors []*ColorName

func (c Colors) Len() int {
	return len(c)
}
func (c Colors) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type ByColor struct {
	Colors
}

func (b ByColor) Less(i, j int) bool {
	return b.Colors[i].Color < b.Colors[j].Color
}

func colorNames() []*ColorName {
	s := make([]*ColorName, len(tcell.ColorNames))
	index := 0
	for name, color := range tcell.ColorNames {
		s[index] = &ColorName{
			Name:  name,
			Color: color,
		}
		index++
	}
	sort.Sort(ByColor{s})
	return s
}

func rgb(c tcell.Color) (uint8, uint8, uint8) {
	r, g, b := c.RGB()
	return uint8(r), uint8(g), uint8(b)
}

func main() {
	if len(os.Args) < 2 {
		gchalk.SetLevel(gchalk.LevelAnsi16m)
	}
	colors := colorNames()
	for _, color := range colors {
		fmt.Println(gchalk.RGB(rgb(color.Color))("■"), color.Name)
	}
}
