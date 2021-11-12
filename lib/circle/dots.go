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

func CreateDots(img image.Image) *Dots {
	bounds := img.Bounds()

	// points
	points := make([]Point, 0)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y)
			_, _, _, a := pixel.RGBA()
			if a != 0 {
				points = append(points, Point{X: x, Y: y})
			}
		}
	}
	numOfPoints := len(points)

	// centroid
	xg := 0
	yg := 0
	for _, p := range points {
		xg += p.X
		yg += p.Y
	}
	centerX := xg / numOfPoints
	centerY := yg / numOfPoints
	pointCentroid := Point{X: centerX, Y: centerY}

	// distances from centroid
	var sum_d float64 = 0
	var minD float64 = math.MaxFloat64
	var maxD float64 = 0
	for _, p := range points {
		d := distanceBetween(pointCentroid, p)
		p.Distance = d
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
