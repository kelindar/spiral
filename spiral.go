package spiral

import (
	"math"
	"math/rand"
)

const twoPi = 2 * math.Pi

// https://arxiv.org/ftp/arxiv/papers/0908/0908.0892.pdf
func galaxy(phi, scale, b, n float64) float64 {
	return scale / (math.Log(b * math.Tan(phi/(2*n))))
}

type Galaxy struct {
}

func (g Galaxy) Generate(random *rand.Rand) {

}

// Arm represents a single arm for the spiral galaxy
type Arm struct {
	Bulge float64 // Greater buldge (B) results in greater arm sweep and smaller bar/bulge
	Tight float64 // Tight controls the tightness (N)
	Pitch float64 // Pitch of the arm
	Noise float64 // Noise controls the randomness factor applied
}

// Generate generates a single spiral galaxy arm. This uses the function described in
// the paper https://arxiv.org/ftp/arxiv/papers/0908/0908.0892.pdf for realism.
func (a Arm) Generate(random *rand.Rand, count int) (stars []Star) {
	for phi := .0; phi < twoPi; phi += twoPi / float64(count) {
		r := 1 / (math.Log(a.Bulge * math.Tan(phi/(2*a.Tight))))
		rx := random.Float64() * a.Noise
		ry := random.Float64() * a.Noise
		stars = append(stars, Star{
			X: r*math.Cos(phi+a.Pitch) + rx,
			Y: r*math.Sin(phi+a.Pitch) + ry,
		})
	}
	return
}

// Star represents a single star
type Star struct {
	X, Y float64
}
