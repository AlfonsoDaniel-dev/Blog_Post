package model2

import "github.com/golang-jwt/jwt"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
