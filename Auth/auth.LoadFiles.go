package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

func LoadFiles(privateBytes, publicBytes string) error {
	var err error

	once.Do(func() {
		err = loadFiles(privateBytes, publicBytes)
	})

	return err
}

func loadFiles(privPath, publicPath string) error {
	privateBytes, err := ioutil.ReadFile(publicPath)
	if err != nil {
		return err
	}
	publicBytes, err := ioutil.ReadFile(privPath)
	if err != nil {
		return err
	}

	return parseRSA(privateBytes, publicBytes)
}

func parseRSA(privatebytes, publicBytes []byte) error {
	var err error

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privatebytes)
	if err != nil {
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}

	return nil
}
