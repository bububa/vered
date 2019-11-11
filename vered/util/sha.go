package util

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"

	uuid "github.com/nu7hatch/gouuid"
)

func Salt() (string, error) {
	token, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return Sha1(token.String()), nil
}

func Sha1(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1Sum(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

func Sha1Base64(data []byte) string {
	return base64.StdEncoding.EncodeToString(Sha1Sum(data))
}
