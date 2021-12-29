package driver

import (
	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/database"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/boltdb/bolt"
)

func (d *driver) DeleteMany(databaseName, collection string, options []types.DeleteOptions) (types.OperationResult, error) {
	var result types.OperationResult
	var deletedDocuments []types.Document

	err := d.db.Batch(func(tx *bolt.Tx) error {
		bucket, err := getCollectionWithUpsert(tx, databaseName, collection, false)
		if err != nil {
			return logs.Wrap(err, "failed to get collection")
		}

		for _, delete := range options {
			matched, err := matchWithCursor(bucket, delete.Filter, 1)
			if err != nil && err != database.ErrNoDocuments {
				result.Error = err.Error()
				continue
			}

			match := matched[0]
			uuid, ok := match[documentIDFieldName].(string)
			if !ok {
				result.Error = "failed to get document uuid"
				continue
			}

			err = bucket.Delete([]byte(uuid))
			if err != nil {
				result.Error = err.Error()
				continue
			}
			deletedDocuments = append(deletedDocuments, match)
			result.Deleted++
		}

		return err
	})
	result.Data = deletedDocuments
	return result, err
}

func (d *driver) DeleteOne(databaseName, collection string, option types.DeleteOptions) (types.OperationResult, error) {
	return d.DeleteMany(databaseName, collection, []types.DeleteOptions{option})
}
