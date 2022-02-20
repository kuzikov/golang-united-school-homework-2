package square

import "testing"

func TestCalcSquare(t *testing.T) {
	type variant struct {
		sideLen float64
		sideNum int
		expect  float64
	}

	tests := []variant{
		// circle
		{10, 0, 314.1592653589793},
		// triangle
		{4, 3, 6.928203230275509},
		// square
		{5, 4, 25},
		// wrong case
		{1, 2, 0},
	}

	for _, v := range tests {
		result := CalcSquare(v.sideLen, Sides(v.sideNum))
		if result != v.expect {
			t.Errorf("want:%v, have:%v", v.expect, result)
		}
	}

}
