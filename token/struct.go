package token

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtMap -
type JwtMap = jwt.MapClaims

// JwtClaims -
type JwtClaims = jwt.Claims

// NewJwtMap -
// sub - 標題
// exp - 過期時間timestamp, 不設置請傳0
// nbf - 生效時間timestamp, 不設置請傳0
func NewJwtMap(sub string, exp int64, nbf int64) jwt.MapClaims {
	j := jwt.MapClaims{
		"iss": issuer,
		"sub": sub,
		"iat": time.Now(),
	}

	if exp != 0 {
		j["exp"] = exp
	}
	if nbf != 0 {
		j["nbf"] = nbf
	}

	return j
}
