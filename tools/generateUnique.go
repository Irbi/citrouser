package tools

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/satori/go.uuid"
	"strings"
)

func GenerateBase64ID(size int) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	encoded := strings.ToLower(base64.URLEncoding.EncodeToString(b))
	return encoded, nil
}

func GenerateUUID() (result string) {
	id := uuid.NewV4()
	return id.String()
}
