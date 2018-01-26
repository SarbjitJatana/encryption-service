package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/SarbjitJatana/encryption-service/client/client"
)

var (
	id        = flag.String("id", "", "The id to store the text against")
	plaintext = flag.String("text", "", "Text to be encrypted")
	key       = flag.String("key", "", "The key used to encrypt the data. Needed for decryption")
)

func main() {
	// Read data entered on the command line
	flag.Parse()

	// Validate inputs
	if len(*id) == 0 {
		fmt.Println("Please supply an id")
	}

	if len(*plaintext) != 0 && len(*key) != 0 {
		fmt.Println("Please supply either text or a decryption key, not both")
	}

	if len(*plaintext) == 0 && len(*key) == 0 {
		fmt.Println("Please supply either text or a decryption key")
	}

	client := client.NewHTTPClient("http://localhost:8000")

	if len(*plaintext) > 0 {
		// Store the value along with the id
		log.Printf("Storing id %s, text: %s", *id, *plaintext)
		key, err := client.Store([]byte(*id), []byte(*plaintext))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error %s", err.Error())
		}
		log.Printf("key: %s", key)
	}

	if len(*key) > 0 {
		// Retrieve the value stored using the key
		log.Printf("Retrieving data using id %s, key: %s", *id, *key)
		plaintext, err := client.Retrieve([]byte(*id), []byte(*key))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error %s", err.Error())
		}
		log.Printf("text: %s", plaintext)
	}
}
