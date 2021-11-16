package circle

import (
	"errors"
	"image"
)

func IsCircle(img image.Image, log func(format string, args ...interface{})) (bool, error) {
	pts := GetPtsFromImage(img)
	dots, err := CreateDots(pts)
	if err != nil {
		return false, err
	}

	avg := dots.DistanceAvarage()
	variance := dots.DistanceVariance()
	std := dots.DistanceSTD()

	log("num of points: %v", dots.numOfPoints)
	log("d_sum: %v", dots.sumOfDistances)
	log("d_avg: %v", avg)
	log("d_min: %v", dots.PointsMin())
	log("d_max: %v", dots.PointsMax())
	log("variance: %v", variance)
	log("STD: %v", std)

	// 判定
	// 標準偏差が大きすぎないか
	if std >= 5 {
		return false, errors.New("variance is too large")
	}
	// 図形が小さすぎないか
	if avg <= 20 {
		return false, errors.New("size of circle is too small")
	}

	return true, nil
}
