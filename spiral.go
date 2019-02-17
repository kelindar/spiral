package spiral

import (
	"math"
)

var epsilon = math.Nextafter(1, 2) - 1

// https://arxiv.org/ftp/arxiv/papers/0908/0908.0892.pdf
func galaxy(phi, scale, b, n float64) float64 {
	return scale / (math.Log(b * math.Tan(phi/(2*n))))
}

func gudermannian(x float64) float64 {
	return 2 * math.Pow(math.Tan(math.Pow(math.E, x)), -1)
}

func lobachevsky(x float64) float64 {
	return 2 * math.Pow(math.Tan(math.Pow(math.E, -x)), -1)
}

func logspiral(r, k, phi float64) float64 {
	return r * math.Pow(math.E, .15*phi)
}

// Polar converts to polar coordinates
func polar(x, y float64) (r float64, phi float64) {
	r = math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	phi = math.Atan2(y, x)
	return
}
