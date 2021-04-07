package geometry

import "math"

func GetCubeVolume(s float64) float64 {
	return math.Pow(s, 3)
}

func GetCubeArea(s float64) float64 {
	return s * s * 6
}
