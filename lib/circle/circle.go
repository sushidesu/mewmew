package circle

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

func IsCircle(base64image string, log func(format string, args ...interface{})) (bool, error) {
	circle, err := base64.StdEncoding.DecodeString(base64image)
	if err != nil {
		log("base64 decode error")
		return false, err
	}
	img, _, err := image.Decode(bytes.NewBuffer(circle))
	if err != nil {
		log("image decode error")
		return false, err
	}

	bounds := img.Bounds()
	log("%v", bounds.Max.X)
	log("%v", bounds.Max.Y)

	newimg := image.NewRGBA(bounds)

	points := make([][2]int, 0)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y)
			_, _, _, a := pixel.RGBA()
			if a != 0 {
				points = append(points, [2]int{x, y})
				newimg.Set(x, y, color.Black)
			} else {
				newimg.Set(x, y, color.White)
			}
		}
	}

	// 中心座標を求める
	num_of_points := len(points)
	xg := 0
	yg := 0
	for _, p := range points {
		x := p[0]
		y := p[1]
		xg += x
		yg += y
	}
	centerX := xg / num_of_points
	centerY := yg / num_of_points

	// 中心を描画
	newimg.Set(centerX, centerY, color.RGBA{
		R: 255,
		B: 0,
		G: 0,
		A: 255,
	})

	// 中心から各点までの距離
	distances := make([]int, 0)
	for _, p := range points {
		x := p[0]
		y := p[1]
		d := distance(centerX, centerY, x, y)
		distances = append(distances, d)
	}

	// 最大/最小/合計
	var maxD float64 = 0
	var minD float64 = math.MaxFloat64
	sum_d := 0
	for _, d := range distances {
		maxD = math.Max(maxD, float64(d))
		minD = math.Min(minD, float64(d))
		sum_d += d
	}
	// 距離の平均
	avg_d := sum_d / num_of_points

	log("points: %v", num_of_points)
	log("d_sum: %v", sum_d)
	log("avg: %v", avg_d)
	log("min: %v", minD)
	log("max: %v", maxD)

	// 分散
	var s_distance float64 = 0
	for _, d := range distances {
		s_distance += math.Pow(float64(d-avg_d), 2)
	}
	s_distance /= float64(num_of_points)
	// 標準偏差
	root_s := math.Sqrt(s_distance)

	log("s: %v", s_distance)
	log("roots: %v", root_s)

	// 最大距離と最大距離を描画
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y)
			_, _, _, a := pixel.RGBA()
			if a != 0 {
				d := distance(centerX, centerY, x, y)
				if float64(d) == minD {
					newimg.Set(x, y, color.RGBA{
						R: 0,
						G: 255,
						B: 0,
						A: 255,
					})
				}
				if float64(d) == maxD {
					newimg.Set(x, y, color.RGBA{
						R: 255,
						G: 0,
						B: 0,
						A: 255,
					})
				}
			}
		}
	}

	// 画像出力
	//	dst, err := os.Create("hoge.png")
	//	if err != nil {
	//		return false, nil
	//	}
	//	err = png.Encode(dst, newimg)
	//	if err != nil {
	//		return false, nil
	//	}

	// 判定
	// 分散が大きすぎないか
	if root_s >= 5 {
		return false, errors.New("variance is too large")
	}
	// 図形が小さすぎないか
	if avg_d <= 20 {
		return false, errors.New("size of circle is too small")
	}

	return true, nil
}

func distance(ax int, ay int, bx int, by int) int {
	val := math.Sqrt(float64((bx-ax)*(bx-ax) + (by-ay)*(by-ay)))
	return int(val)
}
