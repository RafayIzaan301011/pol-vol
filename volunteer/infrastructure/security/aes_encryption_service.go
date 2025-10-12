package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

type AESEncryptionService struct {
	key []byte
}

func NewAESEncryptionService(key []byte) *AESEncryptionService {
	return &AESEncryptionService{key: key}
}

// EncryptMessage encrypts the message using AES-GCM and returns a base64 encoded string.
func (s *AESEncryptionService) Encrypt(message string) (string, error) {
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a random nonce of size required by AES-GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	//  --- extra code to read from file (for testing purposes) ---

	f, err := os.Open("d.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf := make([]byte, 20)
	n, err := f.Read(buf)
	if err != nil {
		return "", err
	}
	println("read bytes: ", n)

	//  ----------------------------------------------------------

	// Seal appends the encrypted data to the nonce and authenticates it
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(message), nil)

	// Encode the ciphertext (nonce + encrypted data) as base64 for storage or transmission
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptMessage decrypts a base64 encoded string encrypted using AES-GCM.
func (s *AESEncryptionService) Decrypt(encodedCiphertext string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(s.key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt and authenticate
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
