package cache2d

import (
	"github.com/redstarcoder/draw2d"
)

var (
	glyphCache map[string]map[rune]*Glyph
//	wordCache map[string]map[string]*Word
)

func init() {
	glyphCache = make(map[string]map[rune]*Glyph)
}

// paintGlyph copies a Glyph from the cache to the gc. If it's not in the cache, it calls renderGlyph first
func paintGlyph(gc draw2d.GraphicContext, x, y float64, chr rune) float64 {
	fontName := gc.GetFontData().Name
	if glyphCache[fontName] == nil {
		glyphCache[fontName] = make(map[rune]*Glyph, 60)
	}
	if glyphCache[fontName][chr] == nil {
		glyphCache[fontName][chr] = renderGlyph(gc, fontName, chr)
	}
	g := glyphCache[fontName][chr].Copy()
	g.Path.Shift(x, y)
	gc.AppendPath(g.Path)
	return g.Width
}

// fillGlyph copies a Glyph from the cache to the gc and fills it. If it's not in the cache, it calls renderGlyph first
func fillGlyph(gc draw2d.GraphicContext, x, y float64, chr rune) float64 {
	fontName := gc.GetFontData().Name
	if glyphCache[fontName] == nil {
		glyphCache[fontName] = make(map[rune]*Glyph, 60)
	}
	if glyphCache[fontName][chr] == nil {
		glyphCache[fontName][chr] = renderGlyph(gc, fontName, chr)
	}
	g := glyphCache[fontName][chr].Copy()
	gc.Save()
	gc.Translate(x, y)
	gc.AppendPath(g.Path)
	gc.Fill()
	gc.Restore()
	return g.Width
}

// renderGlyph renders a Glyph then caches and returns it
func renderGlyph(gc draw2d.GraphicContext, fontName string, chr rune) *Glyph {
	gc.Save()
	defer gc.Restore()
	gc.BeginPath()
	width := gc.CreateStringPath(string(chr), 0, 0)
	return &Glyph{
		Path:  gc.CopyPath(),
		Width: width,
	}
}

func CreateStringPathByGlyph(gc draw2d.GraphicContext, str string, x, y float64) float64 {
	xorig := x
	for _, r := range str {
		x += paintGlyph(gc, x, y, r)
	}
	return x - xorig
}

// FillStringByGlyph uses gc.Translate instead of Path.Shift
func FillStringByGlyph(gc draw2d.GraphicContext, str string, x, y float64) float64 {
	gc.BeginPath()
	xorig := x
	for _, r := range str {
		x += fillGlyph(gc, x, y, r)
	}
	return x - xorig
}

type Glyph struct {
	// Path represents a glyph, it is always at (0, 0)
	Path *draw2d.Path
	// Width is the furthest to the right x coordinate
	Width float64
}

func (g *Glyph) Copy() *Glyph {
	ng := &Glyph{
		Path:  g.Path.Copy(),
		Width: g.Width,
	}
	return ng
}
