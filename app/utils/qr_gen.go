package utils

import "github.com/skip2/go-qrcode"

func GenerateQR(text string) ([]byte, error) {
	png, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil

}
