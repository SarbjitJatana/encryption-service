package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// AESCrypto implementation of the crypto interface providing AES encrypt/decrypt functionality
type AESCrypto struct{}

// Encrypt uses AES encryption on the supplied plain text.
// The encrypted text is then stored.
// The ciphertext is returned to the user.
// The key that was used to encrypt the data is also returned
// or an error if encryption failed.
func (a AESCrypto) Encrypt(plaintext []byte) ([]byte, []byte, error) {
	key := generateKey()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	base64encoded := base64.StdEncoding.EncodeToString(plaintext)

	// Define the final ciphertext here so we can create
	// the initialisation vector directly into the slice
	ciphertext := make([]byte, aes.BlockSize+len(base64encoded))

	// Create initialisation vector
	ivector := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, ivector); err != nil {
		return nil, nil, err
	}

	// Encrypt with cipher feedback mode using the given block
	cfb := cipher.NewCFBEncrypter(block, ivector)
	// XOR each byte of the given slice with a byte from the
	// cipher's key stream
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(base64encoded))

	return ciphertext, key, nil
}

// Decrypt uses the key that was used to encrypt the ciphertext to
// convert it back to plaintext.
// The plaintext is returned or an error if there was a problem decrypting.
func (a AESCrypto) Decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext is incorrect length")
	}

	// Define the slice for the initialisation vector from the ciphertext
	ivector := ciphertext[:aes.BlockSize]
	// And the same for the encrypted text
	text := ciphertext[aes.BlockSize:]

	// Decrypt the ciphertext with cipher feedback mode
	cfb := cipher.NewCFBDecrypter(block, ivector)
	cfb.XORKeyStream(text, text)

	plaintext, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Generate a 32-bit key
func generateKey() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		// handle it
	}
	return key
}

// NewAESCrypto creates a new AES crypto
func NewAESCrypto() *AESCrypto {
	return &AESCrypto{}
}
