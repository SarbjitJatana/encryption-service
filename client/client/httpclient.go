package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	storePath    = "/store"
	retrievePath = "/retrieve"
)

// HTTPClient is a wrapper containing the server URL
type HTTPClient struct {
	ServerURL string
}

// EncryptRequest defines the data that will be sent to the server for storing
type EncryptRequest struct {
	ID        string
	PlainText string
}

// Common contains data that is used in both an encrypt repsonse
// and a decrypt request
type Common struct {
	ID  string
	Key string
}

// EncryptResponse contains the key used for encryption which
// is sent back from the server on a 'Store' request
type EncryptResponse struct {
	Common
}

// DecryptRequest contains the id plus the key that should be used
// to decrypt the stored data
type DecryptRequest struct {
	Common
}

// DecryptResponse defines the data sent back from the server
// on a retrieval request
type DecryptResponse struct {
	PlainText string
}

// Store packages up the data into a JSON structure and makes a call to the server
// to store the data in encrypted form. The server will send back the encrypted key
// that the data has been stored with, or an error if not able to store.
func (c HTTPClient) Store(id, payload []byte) (aesKey []byte, err error) {
	log.Printf("Storing id %s, text: %s", id, payload)

	// Store the id and payload into a Request
	encryptRequest := &EncryptRequest{
		ID:        string(id),
		PlainText: string(payload),
	}
	encryptRequestJSON, err := json.Marshal(encryptRequest)

	requestURL := fmt.Sprintf("%s%s", c.ServerURL, storePath)

	// Send the 'store' request to the server
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(encryptRequestJSON))
	if err != nil {
		return nil, err
	}

	// Read the data sent back, which should contain the encryption key
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jsEncryptResponse EncryptResponse
	err = json.Unmarshal(body, &jsEncryptResponse)
	if err != nil {
		return nil, err
	}

	return []byte(jsEncryptResponse.Key), nil
}

// Retrieve makes a request to the server to retrieve the stored (and encrypted) data.
// The client sends the original id the data is stored against, plus the key that was used
// to encrypt the data.
func (c HTTPClient) Retrieve(id, aesKey []byte) (payLoad []byte, err error) {
	log.Printf("Retrieving data using id %s, key: %s", id, aesKey)
	// Store the id and the aesKey into a Request
	decryptRequest := DecryptRequest{
		Common: Common{
			ID:  string(id),
			Key: string(aesKey),
		},
	}

	decryptRequestJSON, err := json.Marshal(decryptRequest)

	requestURL := fmt.Sprintf("%s%s", c.ServerURL, retrievePath)

	// Send the 'retrieve' request to the server
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(decryptRequestJSON))
	if err != nil {
		return nil, err
	}

	// Read the data sent back, which should contain the encryption key
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jsDecryptResponse DecryptResponse
	err = json.Unmarshal(body, &jsDecryptResponse)
	if err != nil {
		return nil, err
	}

	return []byte(jsDecryptResponse.PlainText), nil
}

// NewHTTPClient creates a new HTTP client with configured URL
func NewHTTPClient(serverURL string) *HTTPClient {
	return &HTTPClient{
		ServerURL: serverURL,
	}
}
