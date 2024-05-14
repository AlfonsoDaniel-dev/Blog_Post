package authorization

import (
	model2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/golang-jwt/jwt"

	"time"
)

func GenerateToken(data model2.Login) (string, error) {
	claims := model2.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "BlogPost.co",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signkey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
