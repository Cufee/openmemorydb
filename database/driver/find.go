package driver

import (
	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/database"
	"byvko.dev/repo/openmemorydb/types"
	"github.com/boltdb/bolt"
)

func (d *driver) FindMany(databaseName, collection string, filter types.Filter, limit int) (types.OperationResult, error) {
	var result types.OperationResult
	var matchedResults []types.Document

	d.db.View(func(tx *bolt.Tx) error {
		bucket, err := getCollectionWithUpsert(tx, databaseName, collection, false)
		if err != nil {
			result.Error = err.Error()
			return logs.Wrap(err, "failed to get collection")
		}
		matchedResults, err = matchWithCursor(bucket, filter, limit)
		if err != nil && err != database.ErrNoDocuments {
			result.Error = err.Error()
			return logs.Wrap(err, "failed to match documents")
		}
		return nil
	})

	result.Data = matchedResults
	result.Matched = len(matchedResults)
	return result, nil
}

func (d *driver) FindOne(databaseName, collection string, filter types.Filter) (types.OperationResult, error) {
	return d.FindMany(databaseName, collection, filter, 1)
}
