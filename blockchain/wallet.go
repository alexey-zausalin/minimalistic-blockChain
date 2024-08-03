package blockchain

import (
	"crypto/rand"
	"crypto/rsa"
)

type Wallet struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

func NewWallet() (*Wallet, error) {
	privateKey, publicKey, err := GenerateRSAKeys()

	if err != nil {
		return nil, err
	}

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}
