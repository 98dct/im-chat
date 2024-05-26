package ctxdata

import "github.com/golang-jwt/jwt/v4"

const Identity = "dct.com"

func GetJwtToken(secretKey string, iat, seconds int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["expire"] = iat + seconds
	claims["iat"] = iat
	claims[Identity] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
