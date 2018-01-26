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

// Data passed as json structure between client and server
type Data struct {
	ID         string
	Key        string
	CipherText string
	PlainText  string
}

func handleStore(w http.ResponseWriter, r *http.Request) {
	// Retrieve values from the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var jsRequest Data
	err = json.Unmarshal(body, &jsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ciphertext, key, err := crypto.Encrypt([]byte(jsRequest.PlainText))
	//ciphertext, key, err := crypto.Encrypt([]byte(plaintext))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("/store: generated key=%#x", key)
	log.Printf("/store: ciphertext=%#x", ciphertext)

	// Store the ciphertext along with the id (hashed)
	store.StoreValue([]byte(jsRequest.ID), ciphertext)

	// Return the key in the response
	responseData := Data{jsRequest.ID, hex.EncodeToString(key), "", ""}
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

	var jsRequest Data
	err = json.Unmarshal(body, &jsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the id to retrieve the ciphertext
	ciphertext, exists := store.RetrieveValue([]byte(jsRequest.ID))
	if !exists {
		http.Error(w, "No ciphertext found for ID", http.StatusBadRequest)
		return
	}

	// Use the key to decrypt the ciphertext
	key, err := hex.DecodeString(jsRequest.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	plaintext, err := crypto.Decrypt(key, ciphertext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("/retrieve: plaintext=%s", plaintext)

	// Return the plaintext in the response
	responseData := Data{"", "", "", string(plaintext)}
	js, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/store", handleStore)
	http.HandleFunc("/retrieve", handleRetrieve)

	// Start a listener
	log.Println("Listening on localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
