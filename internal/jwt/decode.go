package jwt

import (
	"github.com/golang-jwt/jwt"
)

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
var hmacSampleSecret = []byte("my_secret_key")

func Decode(rawToken string) (Token, error) {

	jwtToken, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		// TODO !!!
		// Don't forget to validate the alg is what you expect:
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//	 return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// }

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	return fromJwtToken(jwtToken), err
}
