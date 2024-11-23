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

func Sanitise(usersURL string) string {
	var res string
	if !strings.HasPrefix(usersURL, "http://") && !strings.HasPrefix(usersURL, "https://") {
		res = "https://" + usersURL
	}
	return res
}
