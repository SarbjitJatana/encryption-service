package client

type HTTPClient struct {
	ServerURL string
}

// Store packages up the data into a JSON structure and makes a call to the server
// to store the data in encrypted form. The server will send back the encrypted key
// that the data has been stored with, or an error if not able to store.
func (c HTTPClient) Store(id, payload []byte) (aesKey []byte, err error) {
	// Store the id and payload into a Request
	// Send the 'store' request to the server
	return nil, nil
}

// Retrieve makes a request to the server to retrieve the stored (and encrypted) data.
// The client sends the original id the data is stored against, plus the key that was used
// to encrypt the data.
func (c HTTPClient) Retrieve(id, aesKey []byte) (payLoad []byte, err error) {
	// Store the id and the aesKey into a Request
	// Send the 'retrieve' request to the server

	return nil, nil
}

// NewHTTPClient uses the URL passed in to create a
func NewHTTPClient(serverURL string) *HTTPClient {
	return &HTTPClient{
		ServerURL: serverURL,
	}
}

// func main() {
// 	// Determine from values passed in whether we are storing or retrieving values

// 	client := NewHTTPClient("http://localhost:8080")

// 	// var client Client
// 	// client := NewHTTPClient(..)
// 	// client := NewGRPCClient(..)
// }
