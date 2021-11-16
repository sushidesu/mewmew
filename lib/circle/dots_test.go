package circle_test

import (
	"bytes"
	"encoding/base64"
	"image"
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
	dots, _ := circle.CreateDots(genPts())

	actual := dots.DistanceAvarage()
	if actual != 3 {
		t.Fatalf("%v", actual)
	}
}

func TestDots_Centroid(t *testing.T) {
	dots, _ := circle.CreateDots(genPts())

	actual := dots.PointCentroid()
	if actual.X != 3 || actual.Y != 3 || actual.Distance != 0 {
		t.Logf("%v", dots.Points)
		t.Fatalf("%v", actual)
	}
}

func TestDots_Min(t *testing.T) {
	dots, _ := circle.CreateDots([]circle.Pt{
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

func TestDots_GetPtsFromImage(t *testing.T) {
	data, _ := base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAYAAABw4pVUAAAAAXNSR0IArs4c6QAAA3lJREFUeF7tmD2ojmEYx39HBmVCPgaD8pEsykghlBUZECVJWQzKJGQyGElMFsMpiTJYWCwWH2ERxYIkEZEkH111q7e3c45reJ7r/N/6P3U6Z7je+7re3+/5P899nzF8SREYk5rGw2AhYjeBhViIGAGxcZwQCxEjIDaOE2IhYgTExnFCLESMgNg4ToiFiBEQG8cJsRAxAmLjOCEWIkZAbBwnxELECIiN44RYiBgBsXGcEAsRIyA2jhNiIWIExMZxQixEjIDYOE6IhYgREBvHCbEQMQJi4zghFiJGQGwcJ8RCxAiIjeOEWIgYAbFxnBALESMgNo4TYiFiBMTGcUIsRIyA2DhOiIWIERAbxwmxkF4ILAJWAR+AFRN0eA486aVzx4uOekKWALuAE8AvYDbwdYjRTOAHMLdjdr0sN4pCVgPbgO3AHGAe8BhYNwmhfcCF9pk7vVDscNFRERKPpLXA/nanPwVmAZGQA8CrKZhsBi4Br1ua3nXIr/Ol1IVsBI4CC4EvwDXgIXAZeAkc+Y+MALYMuA18bu+YECR7KQsJkDeA88DFAYJngN/A8STVEHAKON1+r09+blrK1IXcApYPkXkG7AXuJ4n9E3KvrbUj+blpKVMWEiBPAhsGyKxpiYn3SfaKda4Cn4AtiUdcdt1e6pSFxCNrOCETSZoMzALgEHAQWAxsAu72QrHDRdWFjLdn/832nSeSNIwjtsN72rY4XuSxG/sGhCD5S1lIwNvZdllx2LsO/AQOA8ca2TiZxxXb4q3AbiDeMR/bif09MB940LbHFtIBgbhp4q6Pn5VAHAz/tNP3W+B7E/UIuAK8AM4B8b45O7RD62CcfpdQT8hU3z7SE+eQGQNFS9vfb9r/ruLQOFLXKAsZKdDZYS0kS6qozkKKQGfbWEiWVFGdhRSBzraxkCypojoLKQKdbWMhWVJFdRZSBDrbxkKypIrqLKQIdLaNhWRJFdVZSBHobBsLyZIqqrOQItDZNhaSJVVUZyFFoLNtLCRLqqjOQopAZ9tYSJZUUZ2FFIHOtrGQLKmiOgspAp1tYyFZUkV1FlIEOtvGQrKkiuospAh0to2FZEkV1VlIEehsGwvJkiqqs5Ai0Nk2FpIlVVRnIUWgs20sJEuqqM5CikBn21hIllRRnYUUgc62sZAsqaI6CykCnW1jIVlSRXUWUgQ62+Yv8LZLZeQErxIAAAAASUVORK5CYII=")
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatalf(err.Error())
	}

	actual := circle.GetPtsFromImage(img)
	if len(actual) > 300 {
		t.Fatalf("%v", len(actual))
	}
}
