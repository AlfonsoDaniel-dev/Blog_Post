package authorization

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"

	"io/ioutil"
	"sync"
)

var (
	once      sync.Once
	signkey   *rsa.PrivateKey
	verifykey *rsa.PublicKey
)

func LoadFile(privateFile, publicfile string) error {
	var err error
	once.Do(func() {
		err = readfiles(privateFile, publicfile)
	})
	return err
}

func readfiles(privateFilePath string, publicFilePath string) error {
	privateBytes, err := ioutil.ReadFile(privateFilePath)
	if err != nil {
		return err
	}
	publicbytes, err := ioutil.ReadFile(privateFilePath)
	if err != nil {
		return err
	}

	return parseRSA(privateBytes, publicbytes)
}

func parseRSA(privateBytes, publicBytes []byte) error {
	var err error

	signkey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}

	verifykey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}

	return nil
}
