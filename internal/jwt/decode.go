package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

func Decode(rawToken, alg, secret string) (Token, error) {

	jwtToken, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if alg != "" {
			tokenAlg := token.Header["alg"]
			if tokenAlg != alg {
				return nil, fmt.Errorf("expected %s signing method, got %s", alg, tokenAlg)
			}
		}
		return []byte(secret), nil
	})
	return fromJwtToken(jwtToken), err
}
