package store

import (
	"bytes"
	"testing"
)

func TestStoreValue(t *testing.T) {
	memorystore := NewMemoryStore()
	memorystore.StoreValue([]byte("1"), []byte("ciphertext1"))
	memorystore.StoreValue([]byte("2"), []byte("ciphertext2"))

	if len(memorystore.cipherStore) != 2 {
		t.Errorf("Incorrect number of elements in the cipherstore; expected 2, got %d",
			len(memorystore.cipherStore))
	}
}

func TestRetrieveValue(t *testing.T) {
	memorystore := NewMemoryStore()
	memorystore.StoreValue([]byte("1"), []byte("ciphertext1"))
	memorystore.StoreValue([]byte("2"), []byte("ciphertext2"))

	ciphertext, exists := memorystore.RetrieveValue([]byte("1"))
	if !exists {
		t.Errorf("Element not successfully retrieved from cipherstore")
	}

	if !bytes.Equal(ciphertext, []byte("ciphertext1")) {
		t.Errorf("Incorrect value retrieved from cipher store; expected 'ciphertext1', got %s",
			ciphertext)
	}

	ciphertext, exists = memorystore.RetrieveValue([]byte("3"))
	if exists {
		t.Errorf("Unexpected value in the cipherstore %s", ciphertext)
	}
}
