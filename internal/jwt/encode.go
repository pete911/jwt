package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func Encode(claims, alg, key string) (string, error) {

	signingMethod := jwt.GetSigningMethod(alg)
	if signingMethod == nil {
		return "", fmt.Errorf("invalid algorithm %q", alg)
	}

	var mapClaims jwt.MapClaims
	if err := json.Unmarshal([]byte(claims), &mapClaims); err != nil {
		return "", fmt.Errorf("cannot unmarshal claims: %w", err)
	}
	return jwt.NewWithClaims(signingMethod, mapClaims).SignedString([]byte(key))
}
