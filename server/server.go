package main

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SarbjitJatana/encryption-service/server/crypto"
	"github.com/SarbjitJatana/encryption-service/server/store"
)

var (
	memorystore store.DataStore
	aescrypto   crypto.Crypto
)

// Common contains data that is used in both an store repsonse
// and a retrieve request
type Common struct {
	ID  string
	Key string
}

// StoreRequest defines the data that has been sent for storing
type StoreRequest struct {
	ID        string
	PlainText string
}

// StoreResponse contains the key used for encryption which
// is sent back to the client
type StoreResponse struct {
	Common
}

// RetrieveRequest contains the id plus the key that should be used
// to decrypt the stored data
type RetrieveRequest struct {
	Common
}

// RetrieveResponse defines the data sent back to the client
// on a retrieval request which should be the original plain text
type RetrieveResponse struct {
	PlainText string
}

func handleStore(w http.ResponseWriter, r *http.Request) {
	// Retrieve values from the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var jsRequest StoreRequest
	err = json.Unmarshal(body, &jsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Storing text %s against id %s", jsRequest.PlainText, jsRequest.ID)
	ciphertext, key, err := aescrypto.Encrypt([]byte(jsRequest.PlainText))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("/store: generated key=%#x", key)
	log.Printf("/store: ciphertext=%#x", ciphertext)

	// Store the ciphertext along with the id (hashed)
	memorystore.StoreValue([]byte(jsRequest.ID), ciphertext)

	// Return the key in the response
	responseData := StoreResponse{
		Common: Common{
			jsRequest.ID, hex.EncodeToString(key),
		},
	}
	js, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleRetrieve(w http.ResponseWriter, r *http.Request) {
	// Retrieve the id and key from the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var jsRequest RetrieveRequest
	err = json.Unmarshal(body, &jsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the id to retrieve the ciphertext
	ciphertext, exists := memorystore.RetrieveValue([]byte(jsRequest.ID))
	if !exists {
		http.Error(w, "No ciphertext found for ID", http.StatusNotFound)
		return
	}

	// Use the key to decrypt the ciphertext
	key, err := hex.DecodeString(jsRequest.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	plaintext, err := aescrypto.Decrypt(key, ciphertext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("/retrieve: plaintext=%s", plaintext)

	// Return the plaintext in the response
	responseData := RetrieveResponse{PlainText: string(plaintext)}
	js, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	memorystore = store.NewMemoryStore()
	aescrypto = crypto.NewAESCrypto()

	http.HandleFunc("/store", handleStore)
	http.HandleFunc("/retrieve", handleRetrieve)

	// Start a listener
	log.Println("Listening on localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
