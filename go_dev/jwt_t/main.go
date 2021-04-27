package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf"//私钥
)

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}
func main() {
	maxAge := 60*60*24
	customeClaims := &CustomClaims{
		UserId:         11,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge)*time.Second).Unix(),
			Issuer: "jerry",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customeClaims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("token %v\n", tokenString)

	ret, err := ParseToken(tokenString);
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("userinfo : %#v\n", ret)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
