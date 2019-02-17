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
	for i := .0; i < 6; i++{
		series = append(series, drawImage(i))
	}
	
	fmt.Println("saving image")
	saveImage(t, "output", series...)
}

func drawImage(t float64) (points plotter.XYs){
	const r1, a1 = 0.05, -math.Pi 
	const r2, a2 = 0.85, +math.Pi
	const N = 100.0
	
	for i := .0; i < N; i++ {
		r := i * ((r2 - r1) / N)

		//phi := i * ((a2 - a1) / N)
		//phi := math.Pow(math.E, -math.Pow(epsilon*r, 2))
		//phi := math.Pow(r, 2) * math.Log(r)
		//phi := math.Pow(r, 5)
		//phi := lobachevsky(r)
		phi := gudermannian(r)
		phi += t
	
		//r = logspiral(r, .15, phi)
		r = galaxy(r, 1, 5, 3.5)
		
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