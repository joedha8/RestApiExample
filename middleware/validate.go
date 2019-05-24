package middleware

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

func validateBody(encriptBody string, times string) (string, error) {
	zone := time.Now().UTC().Format("1504")
	secretkey := "ASD83838IW" + zone

	fmt.Println("Time", zone)
	fmt.Println("Encrypt", encriptBody)

	if strings.Replace(times, ":", "", -1) == zone {
		if len(encriptBody) == 64 {
			decodedBody := Decrypt(encriptBody, secretkey)

			fmt.Println("Decoded", string(decodedBody))

			if decodedBody != "" {
				return decodedBody, nil
			} else {
				fmt.Println("Decoded Error")
				return "", errors.New("Decoded not valid")
			}
		} else {
			fmt.Println("Chipertext Error")
			return "", errors.New("Ciphertext block size not valid")
		}
	} else {
		fmt.Println("Request Error")
		return "", errors.New("Request not valid")
	}
}

// Encrypts text with the passphrase
func Encrypt(text string, passphrase string) string {
	salt := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err.Error())
	}

	key, iv := __DeriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)

	}

	pad := __PKCS5Padding([]byte(text), block.BlockSize())
	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(pad))
	ecb.CryptBlocks(encrypted, pad)

	return b64.StdEncoding.EncodeToString([]byte("Salted__" + string(salt) + string(encrypted)))
}

// Decrypts encrypted text with the passphrase
func Decrypt(encrypted string, passphrase string) string {
	ct, _ := b64.StdEncoding.DecodeString(encrypted)

	if len(ct) < 16 || string(ct[:8]) != "Salted__" {
		return ""
	}

	salt := ct[8:16]
	ct = ct[16:]
	key, iv := __DeriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	dst := make([]byte, len(ct))
	cbc.CryptBlocks(dst, ct)

	decrypt, err := __PKCS5Trimming(dst)
	if err != nil {
		return ""
	}

	return decrypt
}

func __PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func __PKCS5Trimming(encrypt []byte) (string, error) {
	padding := encrypt[len(encrypt)-1]
	if len(encrypt) > int(padding) {
		return string(encrypt[:len(encrypt)-int(padding)]), nil
	} else {
		return "", errors.New("Padding Size Greater Than Encrypt")
	}
}

func __DeriveKeyAndIv(passphrase string, salt string) (string, string) {
	salted := ""
	dI := ""

	for len(salted) < 48 {
		md := md5.New()
		md.Write([]byte(dI + passphrase + salt))
		dM := md.Sum(nil)
		dI = string(dM[:16])
		salted = salted + dI
	}

	key := salted[0:32]
	iv := salted[32:48]

	return key, iv
}
