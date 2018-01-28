package store

// DataStore provides functions to store and retrieve encrypted data
type DataStore interface {
	// StoreValue accepts the encrypted text and the id to store it against
	StoreValue(id, ciphertext []byte)

	// RetrieveValue accepts the id which is used to retrieve the encrypted text
	RetrieveValue(id []byte) ([]byte, bool)
}
