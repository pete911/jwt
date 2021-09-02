package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

func Decode(rawToken, alg, key string) (Token, error) {

	jwtToken, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if alg != "" {
			tokenAlg := token.Header["alg"]
			if tokenAlg != alg {
				return nil, fmt.Errorf("token algorithm %s does not match specified %s", tokenAlg, alg)
			}
		}
		return []byte(key), nil
	})
	return fromJwtToken(jwtToken), err
}
