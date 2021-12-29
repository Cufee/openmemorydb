package driver

import (
	"encoding/json"

	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/boltdb/bolt"
)

func (d *driver) InsertMany(databaseName, collection string, documents []types.Document) (types.OperationResult, error) {
	var result types.OperationResult
	var insertedDocuments []types.Document

	err := d.db.Batch(func(tx *bolt.Tx) error {
		bucket, err := getCollectionWithUpsert(tx, databaseName, collection, true)
		if err != nil {
			return logs.Wrap(err, "failed to get collection")
		}

		for _, document := range documents {
			uuid, err := checkAndFixDocumentUUID(bucket, document)
			if err != nil {
				result.Error = err.Error()
				continue
			}
			document[documentIDFieldName] = uuid

			value, err := json.Marshal(document)
			if err != nil {
				result.Error = err.Error()
				continue
			}
			err = bucket.Put([]byte(uuid), value)
			if err != nil {
				result.Error = err.Error()
				continue
			}
			result.Created++
			insertedDocuments = append(insertedDocuments, document)
		}
		return nil
	})
	result.Data = insertedDocuments
	return result, err
}

func (d *driver) InsertOne(databaseName, collection string, document types.Document) (types.OperationResult, error) {
	return d.InsertMany(databaseName, collection, []types.Document{document})
}
