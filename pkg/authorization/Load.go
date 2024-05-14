package authorization

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"os"

	"sync"
)

var (
	once      sync.Once
	signkey   *rsa.PrivateKey
	verifykey *rsa.PublicKey
)

func LoadFile(privateFile, publicFile string) error {
	var err error
	once.Do(func() {
		err = readfiles(privateFile, publicFile)
	})
	return err
}

func readfiles(privateFilePath, publicFilePath string) error {
	privateBytes, err := os.ReadFile(privateFilePath)
	if err != nil {
		return err
	}
	publicbytes, err := os.ReadFile(publicFilePath)
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
