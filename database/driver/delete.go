package driver

import (
	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/boltdb/bolt"
)

func (d *driver) DeleteMany(databaseName, collection string, filter types.Filter, limit int) (types.OperationResult, error) {
	var result types.OperationResult
	var deletedDocuments []types.Document

	err := d.db.Batch(func(tx *bolt.Tx) error {
		bucket, err := getCollectionWithUpsert(tx, databaseName, collection, false)
		if err != nil {
			result.Error = err.Error()
			return logs.Wrap(err, "failed to get collection")
		}

		matched, err := matchWithCursor(bucket, filter, limit)
		if err != nil {
			result.Error = err.Error()
			return err
		}

		for _, match := range matched {
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

		return nil
	})
	result.Data = deletedDocuments
	return result, err
}

func (d *driver) DeleteOne(databaseName, collection string, filter types.Filter) (types.OperationResult, error) {
	return d.DeleteMany(databaseName, collection, filter, 1)
}
