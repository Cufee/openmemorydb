package operations

import (
	"errors"

	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/database/driver"
	"byvko.dev/repo/openmemorydb/types"
)

func DeleteMany(request types.OperationRequestDeleteMany) (*types.OperationResult, error) {
	d, err := driver.GetDriver()
	if err != nil {
		return nil, logs.Wrap(err, "failed to get driver")
	}

	result, err := d.DeleteMany(request.Database, request.Collection, request.Filter, request.Limit)
	if err != nil {
		return &result, logs.Wrap(err, "failed to delete document")
	}
	if result.Error != "" {
		return &result, errors.New(result.Error)
	}
	return &result, nil
}

func DeleteOne(request types.OperationRequestDeleteOne) (*types.OperationResult, error) {
	var requestMany types.OperationRequestDeleteMany
	requestMany.Database = request.Database
	requestMany.Collection = request.Collection
	requestMany.Filter = request.Filter
	requestMany.Limit = 1
	return DeleteMany(requestMany)
}
