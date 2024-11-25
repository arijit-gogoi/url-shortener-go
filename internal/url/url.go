package url

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func Shorten(usersURL string) string {
	h := sha256.New()
	h.Write([]byte(usersURL))
	hash := hex.EncodeToString(h.Sum(nil))
	shortURL := hash[:10]
	return shortURL
}

func Sanitise(originalURL string) string {
	if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
		originalURL = "https://" + originalURL
	}
	return originalURL
}
