package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type JsonField map[string]interface{}

func (j JsonField) ToJson() string {
	b, _ := json.Marshal(j)
	return string(b)
}

func (j JsonField) ToJsonIndent() string {
	b, _ := json.MarshalIndent(j, "", "  ")
	return string(b)
}

type Token struct {
	Valid  bool
	Header JsonField
	Claims JsonField
}

func (t Token) String() string {
	return fmt.Sprintf("valid: %t, header: %s, claims: %s", t.Valid, t.Header.ToJson(), t.Claims.ToJson())
}

func (t Token) StringIndent() string {
	return fmt.Sprintf("valid: %t\nheader: %s\nclaims: %s", t.Valid, t.Header.ToJsonIndent(), t.Claims.ToJsonIndent())
}

func fromJwtToken(jwtToken *jwt.Token) Token {

	token := Token{Valid: jwtToken.Valid, Header: jwtToken.Header}
	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok {
		token.Claims = JsonField(claims)
	}
	return token
}
