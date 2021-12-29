package driver

import (
	"encoding/json"
	"fmt"

	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/database"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
)

func getCollectionWithUpsert(tx *bolt.Tx, databaseName, collectionName string, withUpsert bool) (*bolt.Bucket, error) {
	db, err := getDatabaseWithUpsert(tx, databaseName, withUpsert)
	if err != nil {
		return nil, err
	}

	collection, err := getCollectionFromDatabaseWithUpsert(db, collectionName, withUpsert)
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func getDatabaseWithUpsert(tx *bolt.Tx, databaseName string, withUpsert bool) (*bolt.Bucket, error) {
	var err error
	bucket := tx.Bucket([]byte(databaseName))
	if bucket == nil && withUpsert {
		bucket, err = tx.CreateBucketIfNotExists([]byte(databaseName))
		if err != nil {
			return nil, logs.Wrap(err, "failed to create database")
		}
	}
	if bucket == nil {
		return nil, database.ErrDatabaseNotFound
	}
	return bucket, nil
}

func getCollectionFromDatabaseWithUpsert(tx *bolt.Bucket, collectionName string, withUpsert bool) (*bolt.Bucket, error) {
	var err error
	bucket := tx.Bucket([]byte(collectionName))
	if bucket == nil && withUpsert {
		bucket, err = tx.CreateBucketIfNotExists([]byte(collectionName))
		if err != nil {
			return nil, logs.Wrap(err, "failed to create collection")
		}
	}
	if bucket == nil {
		return nil, database.ErrCollectionNotFound
	}
	return bucket, nil
}

func matchWithCursor(bucket *bolt.Bucket, filter map[string]interface{}, limit int) ([]types.Document, error) {
	var result []types.Document
	c := bucket.Cursor()

	for k, v := c.First(); k != nil; k, v = c.Next() {
		var document types.Document = make(types.Document)
		err := json.Unmarshal(v, &document)
		if err != nil {
			return nil, logs.Wrap(err, "failed to unmarshal document")
		}

		matched := database.CompareMaps(filter, document)
		if matched {
			logs.Debug("matched document: %v", document)
			logs.Debug("matched filter: %v", filter)
			result = append(result, document)
		}
		if limit > 0 && len(result) >= limit {
			return result, nil
		}
	}

	if len(result) == 0 {
		return nil, database.ErrNoDocuments
	}
	return result, nil
}

func getNextBucketUUID(b *bolt.Bucket) (string, error) {
	id, err := b.NextSequence()
	if err != nil {
		return "", logs.Wrap(err, "failed to get next sequence")
	}

	uuidTail := uuid.NewString()
	return fmt.Sprintf("%v-%s", id, uuidTail), nil
}

func checkAndFixDocumentUUID(bucket *bolt.Bucket, document types.Document) (string, error) {
	uuid, ok := document[documentIDFieldName].(string)
	if !ok {
		id, err := getNextBucketUUID(bucket)
		uuid = id
		return id, err
	}
	return uuid, nil
}
