package util

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
)

func ConvertBase64ToImage(base64string string) (image.Image, error) {
	b, err := base64.StdEncoding.DecodeString(base64string)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	return img, nil
}

func ConvertImageToBase64(img image.Image) (string, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
