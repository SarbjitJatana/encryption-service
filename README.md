# Encryption Service

A server that encrypts plain text and decrypts ciphertext when supplied with a valid id and an encryption key.
A command-line client is also provided for testing.

Notes:
- Messages are sent betweeen the client and the server using HTTP.
- The datastore is implemented as a in-memory cache.
- Data is sent packaged up in a JSON data structure. However if there is an error, a HTTP error response is sent containing a status code - which the client has to additionally check for. If time allowed, this would have also been sent as an error structure in a JSON object.
- I would look to interpret the error messages produced from the server side encryption and decryption process to give a better idea of where the error is occurring.
- As this is my first Go program, I haven't used external libraries so logging uses the standard Go logger to output to stdout. 

## Testing
Some tests have been written for the memorystore and AES encryption.

Code is split into two packages: <code>client</code> and <code>server</code> which will need to be built and run from these folders. This should produce two binaries: main and server.

The server can be built and started from the server directory:

    go build server.go
    ./server

Once the server is running, build and run the client from client/ :

    go build main.go

    ./main -id=1 -text="Hello"
    ./main -id=1 -key=[key from previous command]

Sample output of the encryption service running:

    $ ./main -id=1 -text="Hello there"
    key: ab4c53040e5b046c29f97ba61dd69d2055162e012f9091a96fdb194c688059da
    $ ./main -id=1 -key=ab4c53040e5b046c29f97ba61dd69d2055162e012f9091a96fdb194c688059da
    text: Hello there
