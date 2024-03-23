package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type Claims struct {
    UUID string `json:"uuid"`
    jwt.RegisteredClaims
}


func ReadPrivateKeyFromFile(filepath string) *rsa.PrivateKey {
	pemBytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("Gagal Membaca Private Key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return privateKey
}

// GenerateJWT generates a JWT token signed with RSA-256
func GenerateJWT(uuid string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		UUID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	privateKey := ReadPrivateKeyFromFile("keys/private.key")

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)

	return tokenString, err
}

func ReadPublicKeyFromFile(filepath string) *rsa.PublicKey {
	pemBytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error File public.key tidak ada: %v", err)
	}

	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("Gagal Decode Public Key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatalf("Error Gagal Parsing: %v", err)
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		log.Fatal("Bukan RSA Public Key")
	}

	return publicKey
}


func VerifyToken(tokenString string) (*Claims, error) {

	publicKey := ReadPublicKeyFromFile("keys/public.key")


	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("GAGAL TIDAK BISA DECODE JWT: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}