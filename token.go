package main

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

func loadPrivateKey(keyPath string) (*rsa.PrivateKey, error) {
	secretKeyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	secretKey, err := jwt.ParseRSAPrivateKeyFromPEM(secretKeyBytes)
	if err != nil {
		return nil, err
	}

	return secretKey, nil
}

func generateToken(userId string) (string, error) {
	secretToken := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			"userId": userId,
		})

	secretTokenString, err := secretToken.SignedString(secretPrivateKey)

	return secretTokenString, err
}
