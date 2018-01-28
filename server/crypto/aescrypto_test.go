package crypto

import (
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	aescrypto := NewAESCrypto()
	plaintext := "cipher this text"
	ciphertext, key, _ := aescrypto.Encrypt([]byte(plaintext))

	uncipheredtext, _ := aescrypto.Decrypt(key, ciphertext)

	if string(uncipheredtext) != plaintext {
		t.Errorf("Unciphered text is not the same as plaintext; Expected %s, got %s",
			plaintext, uncipheredtext)
	}
}

func TestEncryptSuccessDecryptFail(t *testing.T) {
	aescrypto := NewAESCrypto()
	plaintext := "cipher this text"
	ciphertext, _, _ := aescrypto.Encrypt([]byte(plaintext))

	_, err := aescrypto.Decrypt([]byte("1234567"), ciphertext)
	if err == nil {
		t.Error("Expected an error when decrypting using the wrong key")
	}
}
