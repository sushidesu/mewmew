package circle

type Point struct {
	X        int
	Y        int
	Distance float64
}

type Dots struct {
	Points []Point

	pointCentroid Point
	pointMax      Point
	pointMin      Point
}

func (d *Dots) DistanceAvarage() float64 {
	return 100
}

func (d *Dots) DistanceVariance() float64 {
	return 100
}

func (d *Dots) DistanceSTD() float64 {
	return 100
}

func (d *Dots) PointMin() Point {
	return d.pointMin
}

func (d *Dots) PointMax() Point {
	return d.pointMax
}

func (d *Dots) PointCentroid() Point {
	return d.pointCentroid
}

func CreateDots() *Dots {
	return &Dots{
		Points: []Point{},
	}
}
