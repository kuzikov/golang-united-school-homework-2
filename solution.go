package square

import "math"

// Define custom int type to hold sides number.
type Sides int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const (
	SidesCircle   Sides = 0
	SidesTriangle Sides = 3
	SidesSquare   Sides = 4
)

// CalcSquare (10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	switch sidesNum {
	case SidesCircle:
		return (math.Pi * sideLen * sideLen)
	case SidesTriangle:
		return (sideLen * sideLen * math.Sqrt(3) / 4)
	case SidesSquare:
		return (sideLen * sideLen)
	default:
		return 0
	}

}
