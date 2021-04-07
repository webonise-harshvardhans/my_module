package geometry

import "math"

func GetSphereVolume(r float64) float64 {
	return (4 / 3) * math.Pi * r * r
}

func GetSphereArea(r float64) float64 {
	return 4 * math.Pi * r * r
}
