package spiral

import (
	"math"
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
	const arms = 7.0
	for i := .0; i < arms; i++{
		t := i / arms * 2*math.Pi
		series = append(series, drawImage(t))
	}
	
	fmt.Println("saving image")
	saveImage(t, "output", series...)
}

func drawImage(t float64) (points plotter.XYs){
	arm := Arm{ Bulge: 0.55, Tight: 3.7, Pitch: t, Noise: .15}
	for _, v := range arm.Generate(random, 100){
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
	err := p.Save(5*vg.Inch, 5*vg.Inch, name +".png"); 
	assert.NoError(t, err)
}