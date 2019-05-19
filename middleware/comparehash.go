package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func compareHashRequest(key string, hashBody string, time string) string {
	if encodeWithHmac(key, fmt.Sprint(`{"ping_data":"ping"}`, time)) == hashBody {
		return fmt.Sprint(`{"ping_data":"ping"}`)
	}
	return ""
}

func encodeWithHmac(key string, body string) string {
	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(body))

	hash := fmt.Sprintf("%x", hasher.Sum(nil))
	return hash
}
