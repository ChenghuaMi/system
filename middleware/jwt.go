package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

var mySigin = []byte("mch")

type AdminClain struct {
	Id int
	jwt.StandardClaims
}
// 生成token
func CreateToken(id int) (string,error) {
	clain := AdminClain{
		Id:             id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,clain)
	tokenString,err := token.SignedString(mySigin)
	return tokenString,err
}

func ParseToken(tokenString string) (*AdminClain,error){
	token,err := jwt.ParseWithClaims(tokenString,&AdminClain{},func(token *jwt.Token) (interface{}, error) {
		return mySigin,nil
	})
	if clain,ok := token.Claims.(*AdminClain);ok && token.Valid {
		fmt.Println(clain.Id,clain.ExpiresAt)
		return clain,err
	}
	return nil,errors.New("parse token error")

}