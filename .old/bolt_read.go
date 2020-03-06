package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const BLOCKCHAIN_BUCKET = "blockchain"

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKCHAIN_BUCKET))
		if bucket == nil {
			log.Panic("View failed!")
		}
		res := bucket.Get([]byte("name"))
		fmt.Printf("%s", res)
		return nil
	})
}
