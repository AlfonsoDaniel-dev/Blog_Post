package authorization

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

func GetEmailFromJWT(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid token")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", errors.New("cannot decode payload from token")
	}

	var Claims map[string]interface{}
	if err := json.Unmarshal(payload, &Claims); err != nil {
		return "", errors.New("cannot unmarshal payload from token")
	}

	email, ok := Claims["email"].(string)
	if !ok {
		return "", errors.New("cannot find email in token")
	}

	return email, nil
}
