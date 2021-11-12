package circle

import (
	"image"
	"math"
)

type Point struct {
	X        int
	Y        int
	Distance float64
}

type Pt struct {
	X int
	Y int
}

type Dots struct {
	Points      []Point
	numOfPoints int

	sumOfDistances float64

	pointsMax     []Point
	pointsMin     []Point
	pointCentroid Point
}

func (d *Dots) DistanceAvarage() float64 {
	return d.sumOfDistances / float64(d.numOfPoints)
}

func (d *Dots) DistanceVariance() float64 {
	var variance float64 = 0
	for _, p := range d.Points {
		variance += math.Pow(p.Distance-d.DistanceAvarage(), 2)
	}
	return variance / float64(d.numOfPoints)
}

func (d *Dots) DistanceSTD() float64 {
	return math.Sqrt(d.DistanceVariance())
}

func (d *Dots) PointsMin() []Point {
	return d.pointsMin
}

func (d *Dots) PointsMax() []Point {
	return d.pointsMin
}

func (d *Dots) PointCentroid() Point {
	return d.pointCentroid
}

func GetPtsFromImage(img image.Image) []Pt {
	bounds := img.Bounds()

	pts := make([]Pt, 0)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y)
			_, _, _, a := pixel.RGBA()
			if a != 0 {
				pts = append(pts, Pt{X: x, Y: y})
			}
		}
	}
	return pts
}

func CreateDots(pts []Pt) *Dots {
	// points
	numOfPoints := len(pts)
	points := make([]Point, numOfPoints)
	for i, p := range pts {
		points[i] = Point{
			X: p.X,
			Y: p.Y,
		}
	}

	// centroid
	xg := 0
	yg := 0
	for _, p := range pts {
		xg += p.X
		yg += p.Y
	}
	centerX := xg / numOfPoints
	centerY := yg / numOfPoints
	pointCentroid := Point{X: centerX, Y: centerY, Distance: 0}

	// distances from centroid
	var sum_d float64 = 0
	var minD float64 = math.MaxFloat64
	var maxD float64 = 0
	for i, p := range points {
		d := distanceBetween(pointCentroid, p)
		points[i].Distance = d
		sum_d += d
		minD = math.Min(minD, float64(d))
		maxD = math.Max(maxD, float64(d))
	}

	// min/max
	pointsMin := make([]Point, 0)
	pointsMax := make([]Point, 0)
	for _, p := range points {
		if p.Distance == minD {
			pointsMin = append(pointsMin, p)
		}
		if p.Distance == maxD {
			pointsMax = append(pointsMax, p)
		}
	}

	return &Dots{
		Points:      points,
		numOfPoints: numOfPoints,

		sumOfDistances: sum_d,

		pointsMin:     pointsMin,
		pointsMax:     pointsMax,
		pointCentroid: pointCentroid,
	}
}

func distanceBetween(a Point, b Point) float64 {
	val := math.Sqrt(math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2))
	return val
}
