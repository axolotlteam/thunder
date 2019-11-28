package token

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	key    []byte
	method = jwt.SigningMethodHS512
	issuer string
)

// -
const (
	UserClaimsType = "Claims"
	UserMapsType   = "Maps"
)

// JWT -
type JWT interface {
	Create(body interface{}) JWT
	Parse(token string, body interface{}) JWT
	Error() error
	Valid() error
	Get() string
}

type jt struct {
	t   *jwt.Token
	jwt string
	err error
}

// NewJWT -
func NewJWT() JWT {
	return new(jt)
}

// SetKey -
func SetKey(skey string) {
	key = []byte(skey)
}

// SetIssuer -
func SetIssuer(iss string) {
	issuer = iss
}

func (j *jt) Create(body interface{}) JWT {
	switch body.(type) {
	case jwt.MapClaims:
		j.t = jwt.NewWithClaims(method, body.(jwt.MapClaims))
	case jwt.Claims:
		j.t = jwt.NewWithClaims(method, body.(jwt.Claims))
	default:
		j.err = fmt.Errorf("wrong data type , need (JwtMap / JwtClaims)")
	}
	j.jwt, j.err = j.t.SignedString(key)
	return j
}

func (j *jt) Parse(token string, body interface{}) JWT {
	j.t, j.err = jwt.ParseWithClaims(token, body.(jwt.Claims), keyLookup)

	if j.err == nil {
		j.err = j.Valid()
	}

	return j
}

func (j *jt) Valid() error {
	switch j.t.Claims.(type) {
	case jwt.MapClaims:
		t := j.t.Claims.(jwt.MapClaims)
		if !t.VerifyIssuer(issuer, true) {
			return fmt.Errorf("the issuer is not  %v", issuer)
		}
		if t.Valid() != nil {
			return fmt.Errorf("valid token failed")
		}

	case jwt.Claims:
		if j.t.Claims.Valid() != nil {
			return fmt.Errorf("valid token failed")
		}
	default:
		return fmt.Errorf("valid token failed")
	}

	return nil
}

func (j *jt) Error() error {
	return j.err
}

func (j *jt) Get() string {
	return j.jwt
}

func keyLookup(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("jwt signature method faile")
	}
	return key, nil
}
