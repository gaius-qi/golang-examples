package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// SigningMethodRS256
	// New token
	privateKeyPEM, err := ioutil.ReadFile("./private.pem")
	if err != nil {
		panic("could not read signing key")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		fmt.Println(err)
		panic("could not parse signing key")
	}

	claims := &jwt.StandardClaims{}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	stoken, err := token.SignedString(privateKey)
	if err != nil {
		panic("could not sign token using key")
	}

	// Parse token
	publicKeyPEM, err := ioutil.ReadFile("./public.pem")
	if err != nil {
		panic("could not read validation certificate")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		panic("could not parse validation certificate")
	}
	if _, err := jwt.ParseWithClaims(stoken, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unsupported signing method in token")
		}

		return publicKey, nil
	}); err != nil {
		panic("could not parse token")
	}

	fmt.Printf("parse successfully, claims is %v", claims)
}
