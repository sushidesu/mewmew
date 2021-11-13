package circle_test

import (
	"testing"

	"github.com/sushidesu/mewmew/lib/circle"
)

func genPts() []circle.Pt {
	return []circle.Pt{
		{X: 3, Y: 0},
		{X: 6, Y: 3},
		{X: 3, Y: 6},
		{X: 0, Y: 3},
	}
}

func TestDots_DistanceAvarage(t *testing.T) {
	dots := circle.CreateDots(genPts())

	actual := dots.DistanceAvarage()
	if actual != 3 {
		t.Fatalf("%v", actual)
	}
}

func TestDots_Centroid(t *testing.T) {
	dots := circle.CreateDots(genPts())

	actual := dots.PointCentroid()
	if actual.X != 3 || actual.Y != 3 || actual.Distance != 0 {
		t.Logf("%v", dots.Points)
		t.Fatalf("%v", actual)
	}
}

func TestDots_Min(t *testing.T) {
	dots := circle.CreateDots([]circle.Pt{
		{X: 3, Y: 0},
		{X: 6, Y: 3},
		{X: 3, Y: 6},
		{X: 1, Y: 3}, // Min
	})

	actual := dots.PointsMin()
	if len(actual) != 1 {
		t.Fatalf("%v", actual)
	}
	if actual[0].X != 1 || actual[0].Y != 3 {
		t.Fatalf("%v", actual)
	}
}

//func TestDots_Max(t *testing.T) {
//	dots := circle.CreateDots([]circle.Pt{
//		{X: 0, Y: 0},
//		{X: 10, Y: 0},
//		{X: 5, Y: 10},
//	})
//
//	actual := dots.PointsMax()
//	if len(actual) != 1 {
//		t.Logf("%v", dots.Points)
//		t.Fatalf("actual.len != 1: %v", actual)
//	}
//	if actual[0].X != 0 || actual[0].Y != 2 {
//		t.Logf("%v", dots.Points)
//		t.Fatalf("%v", actual)
//	}
//}
