package util

import (
	"strings"

	v1 "k8s.io/api/core/v1"
)

const tokenLabel = "token"

func ExtractToken(secret *v1.Secret) string {
	return strings.TrimSpace(string(secret.Data[tokenLabel]))
}