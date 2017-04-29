package main

import (
	"kvstore"
	"os"
)

func main() {
	known_address:="127.0.0.1:300"+os.Args[1]
	reply := new(string)
	kvstore.Client_Transaction(known_address,reply)
}