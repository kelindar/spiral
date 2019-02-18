package spiral

import (
	"math"
	"math/rand"
)

const twoPi = 2 * math.Pi

// Galaxy represents a spiral galaxy generator.
type Galaxy struct {
	Bulge float64 // Greater buldge (B) results in greater arm sweep and smaller bar/bulge
	Tight float64 // Tight controls the tightness (N)
	Error float64 // Error controls the randomness factor applied
}

// Generate creates a set of arms for the galaxy, with the same properties.
func (g Galaxy) Generate(random *rand.Rand, armCount int) (arms []Arm) {
	for phi := .0; phi < twoPi; phi += twoPi / float64(armCount) {
		arms = append(arms, Arm{
			Bulge: g.Bulge,
			Tight: g.Tight,
			Pitch: phi + normal(random, twoPi/float64(armCount)*g.Error),
			Error: g.Error,
		})
	}
	return
}

// Arm represents a single arm for the spiral galaxy
type Arm struct {
	Bulge float64 // Greater buldge (B) results in greater arm sweep and smaller bar/bulge
	Tight float64 // Tight controls the tightness (N)
	Pitch float64 // Pitch of the arm
	Error float64 // Error controls the randomness factor applied
}

// Generate generates a single spiral galaxy arm. This uses the function described in
// the paper https://arxiv.org/ftp/arxiv/papers/0908/0908.0892.pdf for realism.
func (a Arm) Generate(random *rand.Rand, count int) (stars []Star) {
	for phi := .0; phi < twoPi; phi += twoPi / float64(count) {
		r := 1 / (math.Log(a.Bulge * math.Tan(phi/(2*a.Tight))))
		rx := normal(random, a.Error)
		ry := normal(random, a.Error)
		stars = append(stars, Star{
			X: r*math.Cos(phi+a.Pitch) + rx,
			Y: r*math.Sin(phi+a.Pitch) + ry,
		})
	}
	return
}

// Normal generates a floating point number according to a normal distribution.
func normal(random *rand.Rand, stdDev float64) float64 {
	return random.NormFloat64() * stdDev
}

// Star represents a single star
type Star struct {
	X, Y float64
}
