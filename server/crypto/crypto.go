package crypto

// Crypto provides functions to encrypt and decrypt data
type Crypto interface {
	// Encrypt accepts the plaintext to be encrypted.
	Encrypt(plaintext []byte) ([]byte, []byte, error)

	// Decrypt accepts the key to use to decrypt the given ciphertext
	Decrypt(key, ciphertext []byte) ([]byte, error)
}
