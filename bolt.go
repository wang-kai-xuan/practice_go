package main

import (
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

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKCHAIN_BUCKET))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BLOCKCHAIN_BUCKET))
			if err != nil {
				log.Panic("CreateBucket failed!")
			}
		}
		bucket.Put([]byte("name"), []byte("wangkaixuan"))
		return nil
	})
}
