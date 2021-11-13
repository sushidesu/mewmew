package circle

import (
	"image"
	"image/color"
)

func ShowCircle(dots Dots, rectangle image.Rectangle) *image.RGBA {
	img := image.NewRGBA(rectangle)

	// bg
	for y := rectangle.Min.Y; y < rectangle.Max.Y; y++ {
		for x := rectangle.Min.X; x < rectangle.Max.X; x++ {
			img.Set(x, y, color.White)
		}
	}

	// points
	for _, p := range dots.Points {
		img.Set(p.X, p.Y, color.Black)
	}

	// min
	for _, p := range dots.pointsMin {
		img.Set(p.X, p.Y, color.RGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 0,
		})
	}

	// max
	for _, p := range dots.pointsMax {
		img.Set(p.X, p.Y, color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 0,
		})
	}

	// centroid
	centroid := dots.PointCentroid()
	img.Set(centroid.X, centroid.Y, color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 0,
	})

	return img
}
