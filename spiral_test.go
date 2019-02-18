package spiral

import (
	"fmt"
	"math/rand"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"github.com/stretchr/testify/assert"
)

var random = rand.New(rand.NewSource(42))


func TestBuilder(t *testing.T) {
	series := []plotter.XYs{}
	galaxy := Galaxy{ Bulge: 0.55, Tight: 3.7, Error: .05}
	for _, arm := range galaxy.Generate(random, 7){
		series = append(series, drawArm(arm))
	}
	
	fmt.Println("saving image")
	saveImage(t, "output", series...)
}

func drawArm(arm Arm) (points plotter.XYs){
	for _, v := range arm.Generate(random, 500){
		points = append(points, plotter.XY{ X: v.X, Y: v.Y})
	}
	return
}

func saveImage(t *testing.T, name string, series ...plotter.XYs){
	plotutil.DefaultGlyphShapes = []draw.GlyphDrawer{ draw.CrossGlyph{}	}
	p, _ := plot.New()
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	for _, s := range series{
		plotutil.AddScatters(p, s)
	}
	err := p.Save(7*vg.Inch, 7*vg.Inch, name +".png"); 
	assert.NoError(t, err)
}