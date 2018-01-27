package store

import (
	"crypto/sha256"
	"encoding/hex"
)

// store ciphertext
var cipherStore = make(map[string][]byte)

// StoreValue stores the given ciphertext against the id
// Hash the id first for extra security
func StoreValue(id, ciphertext []byte) {
	hashedID := hash(id)
	cipherStore[hashedID] = ciphertext
}

// RetrieveValue uses the given id to retrieve the ciphertext
// from the store if it exists.
// Hashes the id first in order to retrieve the text
func RetrieveValue(id []byte) ([]byte, bool) {
	hashedID := hash(id)
	ciphertext, exists := cipherStore[hashedID]

	return ciphertext, exists
}

// hash hashes the given ID using SHA256 and encodes to a string
func hash(id []byte) string {
	sha256Byte := sha256.Sum256(id)
	sha256String := hex.EncodeToString(sha256Byte[:])
	return sha256String
}
