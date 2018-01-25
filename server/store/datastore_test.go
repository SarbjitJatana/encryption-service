package store

import (
	"bytes"
	"testing"
)

func TestStoreValue(t *testing.T) {
	StoreValue([]byte("1"), []byte("ciphertext1"))
	StoreValue([]byte("2"), []byte("ciphertext2"))

	if len(cipherStore) != 2 {
		t.Errorf("Incorrect number of elements in the cipherstore; expected 2, got %d",
			len(cipherStore))
	}
}

func TestRetrieveValue(t *testing.T) {
	StoreValue([]byte("1"), []byte("ciphertext1"))
	StoreValue([]byte("2"), []byte("ciphertext2"))

	ciphertext, exists := RetrieveValue([]byte("1"))
	if !exists {
		t.Errorf("Element not successfully retrieved from cipherstore")
	}

	if !bytes.Equal(ciphertext, []byte("ciphertext1")) {
		t.Errorf("Incorrect value retrieved from cipher store; expected 'ciphertext1', got %s",
			ciphertext)
	}

	ciphertext, exists = RetrieveValue([]byte("3"))
	if exists {
		t.Errorf("Unexpected value in the cipherstore %s", ciphertext)
	}
}
