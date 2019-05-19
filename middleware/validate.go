package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

const separator = ":"

func validateBody(encriptBody string, time string, keyToCompare string) (string, error) {

	s := strings.Split(encriptBody, separator)

	encriptjsonBody := s[0]
	keyInBody := s[1]

	decodedBody, err := base64Decript(encriptjsonBody)
	if err != nil {
		return "", err
	}

	if keyInBody != encodeWithHmac(keyToCompare, fmt.Sprintf("%s%s", decodedBody, time)) {
		return "", errors.New("key invalid")
	}

	return decodedBody, nil

}

func base64Decript(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func encodeWithHmac(key string, body string) string {
	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(body))

	hash := fmt.Sprintf("%x", hasher.Sum(nil))
	return hash
}
