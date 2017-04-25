package kvstore

import (
	"crypto/sha1"
	// "fmt"
	"hash"
	"time"
	"errors"
)

type Config struct {
	HashFunc      func() hash.Hash // Hash function to use
	StabilizeMin  time.Duration    // Minimum stabilization time
	StabilizeMax  time.Duration    // Maximum stabilization time
	NumSuccessors int              // Number of successors to maintain
	NumReplicas	  int 			   // Number of replicas
	// Delegate      Delegate         // Invoked to handle ring events
	hashBits      int              // Bit size of the hash function
}

// Represents a node, local or remote
type Node struct {
	Id   []byte // Unique Chord Id
	Address string // Host identifier
}

type LocalNode struct {
	Node
	successors  []*Node
	finger      []*Node
	data		[]map[string]string
	// last_finger int
	predecessor *Node
	config Config
	// stabilized  time.Time
	// timer       *time.Timer
}

func DefaultConfig() Config {
	return Config{
		sha1.New, // SHA1
		time.Duration(1 * time.Second),
		time.Duration(3 * time.Second),
		3,   // 3 successors
		3,   // 3 Replicas
		// nil, // No delegate
		16, // 16bit hash function
	}
}

func (ln *LocalNode) Init(config Config) {
	// Generate an ID
	ln.genId()
	ln.config = config
	// Initialize all state
	ln.successors = make([]*Node, ln.config.NumSuccessors)
	ln.finger = make([]*Node, ln.config.hashBits)

	// // Register with the RPC mechanism
	// ln.ring.transport.Register(&ln.Node, ln)
}

func (ln *LocalNode) genId() {
	// Use the hash funciton
	conf := ln.config
	hash := conf.HashFunc()
	hash.Write([]byte(ln.Address))

	// Use the hash as the ID
	ln.Id = hash.Sum(nil)
}

func (ln *LocalNode) split_map(data map[string]string, id []byte) map[string]string{
	var new_map map[string]string
	for key,val := range data{
		if(GenHash(ln.config,key).compare(id)<=0){
			new_map[key] = val
			delete(data,key)
		}
	}
	return new_map	
}
func (ln *LocalNode) add_map(target map[string]string, source map[string]string) error{
	for key,val := range source{
		_ , ok = target[key]
		if ok == true {
			return errors.New("Key already in target map")
		}else{
			target[key] = val
		}
	}
	return nil
}