package boltdb

import (
	"fmt"
	"time"

	v1alpha1 "github.com/KommodoreX/dp-rudder/api/v1alpha1/config"
	"github.com/KommodoreX/dp-rudder/pkg/logger"
	"go.etcd.io/bbolt"
)

// DBPath is the path to the BoltDB file.
const DBPath = "your_database.db"

// BoltDBClient is a struct that holds the BoltDB instance.
type BoltDBClient struct {
	db *bbolt.DB
}

// NewBoltDBClient initializes and returns a new BoltDBClient.
func NewBoltDBClient(config v1alpha1.DataStoreConfig) (*BoltDBClient, error) {
	db, err := bbolt.Open(config.DBPath, 0600, &bbolt.Options{Timeout: time.Duration(time.Duration(config.Timeout).Seconds())})
	if err != nil {
		return nil, err
	}

	return &BoltDBClient{db}, nil
}

// Close closes the BoltDB database.
func (c *BoltDBClient) Close() {
	c.db.Close()
}

// Create inserts a key-value pair into the BoltDB database.
func (c *BoltDBClient) Create(bucketName string, key []byte, value []byte) error {
	return c.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			logger.LoggerRudder.Base().Error(err.Error())
			return err
		}
		return bucket.Put(key, value)
	})
}

// Read retrieves the value for a given key from the BoltDB database.
func (c *BoltDBClient) Read(bucketName string, key []byte) ([]byte, error) {
	var value []byte
	err := c.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}
		value = bucket.Get(key)
		return nil
	})
	if err != nil {
		logger.LoggerRudder.Base().Error(err.Error())
		return nil, err
	}
	return value, nil
}

// Update updates the value for a given key in the BoltDB database.
func (c *BoltDBClient) Update(bucketName string, key []byte, value []byte) error {
	return c.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}
		return bucket.Put(key, value)
	})
}

// Delete removes a key-value pair from the BoltDB database.
func (c *BoltDBClient) Delete(bucketName string, key []byte) error {
	return c.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}
		return bucket.Delete(key)
	})
}
