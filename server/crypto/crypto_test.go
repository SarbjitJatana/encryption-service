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
