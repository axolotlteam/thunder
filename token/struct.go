package token

import (
	"fmt"
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
		"iat": fmt.Sprintf("%v", time.Now().Unix()),
	}

	if exp != 0 {
		j["exp"] = fmt.Sprintf("%v", exp)
	}
	if nbf != 0 {
		j["nbf"] = fmt.Sprintf("%v", nbf)
	}

	return j
}
