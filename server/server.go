package server

import (
	"net/http"
	"strconv"

	"github.com/SarbjitJatana/encryption-service/server/crypto"
)

/*func handleStore(id, payload []byte) (aesKey []byte, err error) {

}*/
func handleStore(w http.ResponseWriter, r *http.Request) {
	// Retrieve values from the request
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	plaintext := r.PostFormValue("text")

	ciphertext, key, err := crypto.Encrypt(plaintext)

	// Store the ciphertext along with the id (hashed)

	// Return the key in the response

}

func handleRetrieve(w http.ResponseWriter, r *http.Request) {
	// Retrieve the id and key from the request

	// Use the id to retrieve the ciphertext

	// Use the key to decrypt the ciphertext

	// Return the plaintext in the response
}

func main() {
	http.HandleFunc("/store", handleStore)
	// handle retrieve

	// Start a listener
}
