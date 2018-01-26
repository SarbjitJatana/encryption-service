package client

type HttpClient struct {
	ServerURL string
}

// Store does blah
func (c HttpClient) Store(id, payload []byte) (aesKey []byte, err error) {
	// Store the id and payload into a Request
	// Send the 'store' request to the server
	return nil, nil
}

func (c HttpClient) Retrieve(id, aesKey []byte) (payLoad []byte, err error) {
	// Store the id and the aesKey into a Request
	// Send the 'retrieve' request to the server

	return nil, nil
}

func NewHttpClient(serverURL string) *HttpClient {
	return &HttpClient{
		ServerURL: serverURL,
	}
}

func main() {
	// Determine from values passed in whether we are storing or retrieving values

	client := NewHttpClient("http://localhost:8080")

	// var client Client
	// client := NewHttpClient(..)
	// client := NewGRPCClient(..)
}
