package authorization

import (
	"errors"
	model2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(t string) (model2.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model2.Claim{}, VerifyFunc)
	if err != nil {
		return model2.Claim{}, err
	}

	if !token.Valid {
		return model2.Claim{}, errors.New("token is not invalid")
	}

	claim, ok := token.Claims.(*model2.Claim)
	if !ok {
		return model2.Claim{}, errors.New("token claim is not valid")
	}

	return *claim, nil
}

func VerifyFunc(t *jwt.Token) (interface{}, error) {
	return signkey, nil
}
