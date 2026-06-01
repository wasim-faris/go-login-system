package helpers

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	User_id string `json: "user_id"`
	Email string `json: "email"`
	Role string `json: "role"`

	jwt.RegisteredClaims
}

var jwtKey []byte

func SetJwtKey(key string){
	jwtKey=[]byte(key)
}

func GetJWTKey()[]byte{
	return jwtKey
}

func ValidateToken(tokenString string)(Claims, error){
	secretkey:=GetJWTKey()

	token,err:= jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token)(interface{}, error){
		return secretkey,nil
	})

	if err!=nil{
	return Claims{},err
}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
    return *claims, nil
}
	return Claims{},errors.New("invalid token")
}

