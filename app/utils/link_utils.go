package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GetURLHash(url string) string {
	hash := sha256.Sum256([]byte(url))
	// Base64 URL-safe без padding
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash[:])[:12]
}
