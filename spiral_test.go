package spiral

import (
	"math"
	"fmt"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	series := []plotter.XYs{}
	const arms = 6.0
	for i := .0; i < arms; i++{
		t := i / arms * 2*math.Pi
		println(t)
		series = append(series, drawImage(t))
	}
	assert.Fail(t, "")
	
	fmt.Println("saving image")
	saveImage(t, "output", series...)
}

func drawImage(t float64) (points plotter.XYs){
	const r1, a1 = 0.05, 0
	const r2, a2 = 0.85, 2*math.Pi
	const N = 100.0
	
	for i := .0; i < N; i++ {
		phi := i * ((a2 - a1) / N)
		r := galaxy(phi, 1, .5, 4)
		phi += t

		x := r * math.Cos(phi)
		y := r * math.Sin(phi)
		points = append(points, plotter.XY{ X: x, Y: y})
	}
	return
}

func saveImage(t *testing.T, name string, series ...plotter.XYs){
	p, _ := plot.New()
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	for _, s := range series{
		plotutil.AddLinePoints(p, "", s)
	}
	err := p.Save(4*vg.Inch, 4*vg.Inch, name +".png"); 
	assert.NoError(t, err)
}