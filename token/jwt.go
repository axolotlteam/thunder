package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/maxjkfc/cocola/errors"
)

var (
	key    []byte
	method = jwt.SigningMethodHS512
	issuer string
)

const (
	UserClaimsType = "Claims"
	UserMapsType   = "Maps"
)

type JWT interface {
	Create(body interface{}) JWT
	Parse(token string, body interface{}) JWT
	Error() errors.Error
	Valid() errors.Error
	Get() string
}

type jt struct {
	t    *jwt.Token
	jwt  string
	err  error
	errs errors.Error
}

func New() JWT {
	return new(jt)
}

func SetKey(skey string) {
	key = []byte(skey)
}

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
		j.err = errors.ErrorJwtCreateFailedForType
	}
	j.jwt, j.err = j.t.SignedString(key)
	return j
}

func (j *jt) Parse(token string, body interface{}) JWT {
	j.t, j.err = jwt.ParseWithClaims(token, body.(jwt.Claims), keyLookup)

	if j.err == nil {
		j.errs = j.Valid()
	} else {
		j.errs = errors.T(10000, j.err.Error())
	}

	return j
}

func (j *jt) Valid() errors.Error {
	switch j.t.Claims.(type) {
	case jwt.MapClaims:
		t := j.t.Claims.(jwt.MapClaims)
		if !t.VerifyIssuer(issuer, true) {
			return errors.ErrorJwtWrongIssuer
		}
		if t.Valid() != nil {
			return errors.ErrorJwtValidFailed
		}

	case jwt.Claims:
		if j.t.Claims.Valid() != nil {
			return errors.ErrorJwtValidFailed
		}
	default:
		return errors.ErrorJwtValidFailed
	}

	return nil
}

func (j *jt) Error() errors.Error {
	return j.errs
}

func (j *jt) Get() string {
	return j.jwt
}

func keyLookup(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.ErrorJwtSigningMethod
	}
	return key, nil
}
