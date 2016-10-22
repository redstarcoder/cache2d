package cache2d

import (
	"fmt"
	"github.com/redstarcoder/draw2d"
)

var (
	glyphCache map[string]map[rune]*Glyph
)

func init() {
	glyphCache = make(map[string]map[rune]*Glyph)
}

// fillGlyph copies a Glyph from the cache to the gc and fills it
func fillGlyph(gc draw2d.GraphicContext, x, y float64, fontName string, chr rune) float64 {
	g := fetchGlyph(gc, fontName, chr)
	gc.Save()
	gc.BeginPath()
	gc.Translate(x, y)
	gc.Fill(g.Path)
	gc.Restore()
	return g.Width
}

// strokeGlyph fetches a Glyph from the cache to the gc and strokes it
func strokeGlyph(gc draw2d.GraphicContext, x, y float64, fontName string, chr rune) float64 {
	g := fetchGlyph(gc, fontName, chr)
	gc.Save()
	gc.BeginPath()
	gc.Translate(x, y)
	gc.Stroke(g.Path)
	gc.Restore()
	return g.Width
}

// fetchGlyph fetches a Glpyh from the cache, calling renderGlyph first if it doesn't already exist
func fetchGlyph(gc draw2d.GraphicContext, fontName string, chr rune) *Glyph {
	if glyphCache[fontName] == nil {
		glyphCache[fontName] = make(map[rune]*Glyph, 60)
	}
	if glyphCache[fontName][chr] == nil {
		glyphCache[fontName][chr] = renderGlyph(gc, fontName, chr)
	}
	return glyphCache[fontName][chr].Copy()
}

// renderGlyph renders a Glyph then caches and returns it
func renderGlyph(gc draw2d.GraphicContext, fontName string, chr rune) *Glyph {
	gc.Save()
	defer gc.Restore()
	gc.BeginPath()
	width := gc.CreateStringPath(string(chr), 0, 0)
	path := gc.GetPath()
	return &Glyph{
		Path:  &path,
		Width: width,
	}
}

// FillStringByGlyph draws a string using glyphs in the cache, rendering them if they don't exist
func FillStringByGlyph(gc draw2d.GraphicContext, str string, x, y float64) float64 {
	xorig := x
	fontData := gc.GetFontData()
	fontName := fmt.Sprintf("%s:%d:%d:%d", fontData.Name, fontData.Family, fontData.Style, gc.GetFontSize())
	for _, r := range str {
		x += fillGlyph(gc, x, y, fontName, r)
	}
	return x - xorig
}

type Glyph struct {
	// Path represents a glyph, it is always at (0, 0)
	Path *draw2d.Path
	// Width is the width of the glyph
	Width float64
}

func (g *Glyph) Copy() *Glyph {
	ng := &Glyph{
		Path:  g.Path.Copy(),
		Width: g.Width,
	}
	return ng
}
