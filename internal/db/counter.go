package db

import (
	"encoding/binary"
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

type CounterDB struct {
	db *bbolt.DB
}

const (
	counterBucket = "counters"
	counterKey    = "main_counter"
)

func NewCounterDB(dbPath string) (*CounterDB, error) {
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create bucket if it doesn't exist
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(counterBucket))
		return err
	})
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create bucket: %w", err)
	}

	return &CounterDB{db: db}, nil
}

func (c *CounterDB) Close() error {
	return c.db.Close()
}

func (c *CounterDB) GetCounter() (int, error) {
	var count int
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(counterBucket))
		v := b.Get([]byte(counterKey))
		if v == nil {
			count = 0
			return nil
		}
		count = int(binary.LittleEndian.Uint64(v))
		return nil
	})
	return count, err
}

func (c *CounterDB) IncrementCounter() (int, error) {
	var newCount int
	err := c.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(counterBucket))
		v := b.Get([]byte(counterKey))

		currentCount := 0
		if v != nil {
			currentCount = int(binary.LittleEndian.Uint64(v))
		}

		newCount = currentCount + 1

		// Convert to bytes
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, uint64(newCount))

		return b.Put([]byte(counterKey), buf)
	})

	if err != nil {
		log.Printf("Error incrementing counter: %v", err)
	}

	return newCount, err
}
