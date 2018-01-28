package main

import (
	"flag"
	"log"
	"os"

	"github.com/SarbjitJatana/encryption-service/client/client"
)

var (
	id        = flag.String("id", "", "The id to store the text against")
	plaintext = flag.String("text", "", "Text to be encrypted")
	key       = flag.String("key", "", "The key used to encrypt the data. Needed for decryption")
)

// Read data entered on the command line and validate it
func validateInputs() {
	flag.Parse()

	if len(*id) == 0 {
		log.Println("Please supply an id")
		os.Exit(1)
	}

	if len(*plaintext) != 0 && len(*key) != 0 {
		log.Println("Please supply either text or a decryption key, not both")
		os.Exit(1)
	}

	if len(*plaintext) == 0 && len(*key) == 0 {
		log.Println("Please supply either text or a decryption key")
		os.Exit(1)
	}
}

func main() {
	validateInputs()

	client := client.NewHTTPClient("http://localhost:8000")

	if len(*plaintext) > 0 {
		// Store the value along with the id
		key, err := client.Store([]byte(*id), []byte(*plaintext))
		if err != nil {
			log.Printf("Error %s", err.Error())
		}
		log.Printf("key: %s", key)
	}

	if len(*key) > 0 {
		// Retrieve the value stored using the key
		plaintext, err := client.Retrieve([]byte(*id), []byte(*key))
		if err != nil {
			log.Printf("Error %s", err.Error())
		}
		log.Printf("text: %s", plaintext)
	}
}
