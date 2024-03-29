package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sushidesu/mewmew/lib/circle"
	"github.com/sushidesu/mewmew/lib/util"
)

type detectionRequest struct {
	Image string `json:"image"`
}

type detectionResponse struct {
	Result   string  `json:"result"`
	Variance float64 `json:"variance"`
	Sd       float64 `json:"sd"`
	Average  float64 `json:"average"`
}

func (r *detectionRequest) validate() error {
	if r.Image == "" {
		return errors.New("image is required")
	}
	return nil
}

func DetectionHandler(w http.ResponseWriter, r *http.Request) {
	// allow CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
		return
	}

	// parse json
	defer r.Body.Close()
	var jsonBody detectionRequest
	err := json.NewDecoder(r.Body).Decode(&jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate
	if err = jsonBody.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// decode image
	img, err := util.ConvertBase64ToImage(jsonBody.Image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate image size
	bounds := img.Bounds()
	if bounds.Max.X > 500 || bounds.Max.Y > 500 {
		http.Error(w, "image size is too large", http.StatusBadRequest)
		return
	}
	// get dots
	dots, err := circle.CreateDots(circle.GetPtsFromImage(img))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// visualize
	result := circle.ShowCircle(*dots, bounds)
	str, err := util.ConvertImageToBase64(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// response
	response := detectionResponse{
		Result:   str,
		Variance: dots.DistanceVariance(),
		Sd:       dots.DistanceSTD(),
		Average:  dots.DistanceAvarage(),
	}

	json.NewEncoder(w).Encode(response)
}
