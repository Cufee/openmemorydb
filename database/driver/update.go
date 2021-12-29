package driver

import (
	"encoding/json"

	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/database"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/boltdb/bolt"
)

func (d *driver) UpdateMany(databaseName, collection string, updates []types.UpdateOptions) (types.OperationResult, error) {
	var result types.OperationResult
	var updatedDocuments []types.Document

	err := d.db.Batch(func(tx *bolt.Tx) error {
		bucket, err := getCollectionWithUpsert(tx, databaseName, collection, false)
		if err != nil {
			return logs.Wrap(err, "failed to get collection")
		}

		for _, update := range updates {
			results, err := matchWithCursor(bucket, update.Filter, 1)
			if err == database.ErrNoDocuments && update.Upsert {
				err = nil // ignore error
			}
			if err != nil {
				result.Error = err.Error()
				continue
			}

			uuid, err := checkAndFixDocumentUUID(bucket, update.Update)
			if err != nil {
				return logs.Wrap(err, "failed to check and fix document uuid")
			}
			update.Update[documentIDFieldName] = uuid

			value, err := json.Marshal(update.Update)
			if err != nil {
				return logs.Wrap(err, "failed to marshal update")
			}

			err = bucket.Put([]byte(uuid), value)
			if err != nil {
				result.Error = err.Error()
				continue
			}
			updatedDocuments = append(updatedDocuments, update.Update)

			if len(results) > 0 {
				result.Updated++
				continue
			}
			result.Created++
		}

		return err
	})
	result.Data = updatedDocuments
	return result, err
}

func (d *driver) UpdateOne(databaseName, collection string, filter, update types.UpdateOptions) (types.OperationResult, error) {
	return d.UpdateMany(databaseName, collection, []types.UpdateOptions{update})
}
