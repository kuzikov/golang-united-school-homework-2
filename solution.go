package square

import "math"

// Define custom int type to hold sides number.
type side int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const (
	circle   side = 0
	triangle side = 3
	square   side = 4
)

// CalcSquare (10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
func CalcSquare(sideLen float64, sidesNum side) float64 {
	switch sidesNum {
	case circle:
		return (math.Pi * sideLen * sideLen)
	case triangle:
		return (sideLen * sideLen * math.Sqrt(3) / 4)
	case square:
		return (sideLen * sideLen)
	default:
		return 0.0
	}

}
