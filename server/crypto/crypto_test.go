package crypto

import (
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	plaintext := "cipher this text"
	ciphertext, key, _ := Encrypt([]byte(plaintext))

	uncipheredtext, _ := Decrypt(key, ciphertext)

	if string(uncipheredtext) != plaintext {
		t.Errorf("Unciphered text is not the same as plaintext; Expected %s, got %s",
			plaintext, uncipheredtext)
	}
}

func TestEncryptSuccessDecryptFail(t *testing.T) {
	plaintext := "cipher this text"
	ciphertext, _, _ := Encrypt([]byte(plaintext))

	_, err := Decrypt([]byte("1234567"), ciphertext)
	if err == nil {
		t.Error("Expected an error when decrypting using the wrong key")
	}
}
