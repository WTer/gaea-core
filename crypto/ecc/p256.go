package ecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"
)

// GenerateECP256Keypair generate and return private and public key byte slices using curve p256
func GenerateECP256Keypair() (privBytes []byte, pubBytes []byte, err error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to generate ecdsa key using curve p256, error: %v", err)
	}

	privBytes, err = x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to marshal EC private key, error: %v", err)
	}

	pubBytes = elliptic.Marshal(elliptic.P256(), priv.X, priv.Y)
	return
}

// FromPubHex convert public key hex string to ecdsa public key
func FromPubHex(pubHex string) (*ecdsa.PublicKey, error) {
	pubBytes, err := hex.DecodeString(pubHex)
	if err != nil {
		return nil, fmt.Errorf("Error getting byte slice from pubHex: %s, error: %v", pubHex, err)
	}

	return FromPubBytes(pubBytes)
}

// FromPrivBytes convert private key bytes to ecdsa private key
func FromPrivBytes(privBytes []byte) (*ecdsa.PrivateKey, error) {
	privKey, err := x509.ParsePKCS8PrivateKey(privBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse pkcs8 private key from private bytes")
	}

	return privKey.(*ecdsa.PrivateKey), nil
}

// FromPubBytes convert public key bytes to ecdsa public key
func FromPubBytes(pubBytes []byte) (*ecdsa.PublicKey, error) {
	x, y := elliptic.Unmarshal(elliptic.P256(), pubBytes)
	if x == nil || y == nil {
		return nil, fmt.Errorf("Error to unmarshal ecdsa publickey, pubBytes: %v", pubBytes)
	}

	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}, nil
}
