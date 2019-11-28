package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Account string   `json:"account"`
	Name    string   `json:"name"`
	Email   []string `json:"email"`
	jwt.StandardClaims
}

func (u *User) Valid() error {
	if u.Issuer != "max" {
		return fmt.Errorf("Error: Wrong Issuer , %s", u.Issuer)
	}
	return nil
}

type UserMap = jwt.MapClaims

var (
	u1 *User
	u2 UserMap
	t1 string
	t2 string
	r1 *User
	r2 UserMap
)

func init() {
	SetKey("mm")
	SetIssuer("max")
	u1 = &User{
		Account: "a1",
		Name:    "a1Name",
		Email: []string{
			"a1@gmail.com",
			"a1@outlook.com",
		},
	}
	u1.Issuer = "max"
	u1.IssuedAt = time.Now().Unix()

	u2 = make(UserMap)

	u2["iss"] = "max"
	u2["account"] = "a2"
	u2["name"] = "a2name"
	u2["email"] = []string{
		"a2@gmail.com",
		"a2@outlook.com",
	}
	u2["ist"] = time.Now().Unix()
}

func Test_CreateJWTWithClaims(t *testing.T) {
	j := NewJWT().Create(u1)
	assert.NoError(t, j.Error())
	spew.Dump(j)
}

func Test_CreateJWTWithMap(t *testing.T) {
	u2["ist"] = time.Now().Unix()
	j := NewJWT().Create(u2)
	assert.NoError(t, j.Error())
	spew.Dump(j)
}

func Test_ParseJWTWithClaims(t *testing.T) {
	r1 = new(User)
	j := NewJWT().Parse(t1, r1)

	if j.Error() != nil {
		t.Errorf("Test_ParseJWTWithClaims Error: %v", j.Error())
	} else {
		t.Logf("Test_ParseJWTWithClaims Sussess: %v", r1)
	}
}

func Test_ParseJWTWithMap(t *testing.T) {
	r2 = make(UserMap)
	j := NewJWT().Parse(t2, r2)
	if j.Error() != nil {
		t.Errorf("Test_ParseJWTWithMap Error: %v", j.Error())
	} else {
		t.Logf("Test_ParseJWTWithMap Sussess: %v", r2)
	}
}

func Benchmark_CreateJWTWithClaims(b *testing.B) {
	j := NewJWT()
	for i := 0; i < b.N; i++ {
		t1 = j.Create(u1).Get()
	}
}

func Benchmark_CreateJWTWithMap(b *testing.B) {
	j := NewJWT()
	u2["ist"] = time.Now().Unix()
	for i := 0; i < b.N; i++ {
		t2 = j.Create(u2).Get()
	}
}

func Benchmark_ParseJWTWithClaims(b *testing.B) {
	j := NewJWT()
	r := &User{}
	for i := 0; i < b.N; i++ {
		j.Parse(t1, r)
	}
}

func Benchmark_ParseJWTWithMap(b *testing.B) {
	j := NewJWT()
	r := make(UserMap)
	for i := 0; i < b.N; i++ {
		j.Parse(t2, r)
	}
}
