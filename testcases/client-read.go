package main

import (
	"kvstore"
	// "strconv"
	"os"
)

func main() {
	key:="key"
	known_address:="127.0.0.1:300"+os.Args[1]
	var result string
	kvstore.Client_remoteRead(known_address,key+os.Args[2],&result)
}