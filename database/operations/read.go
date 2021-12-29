package operations

import (
	"errors"

	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/database/driver"
	"byvko.dev/repo/openmemorydb/types"
)

func ReadMany(request types.OperationRequestReadMany) (*types.OperationResult, error) {
	d, err := driver.GetDriver()
	if err != nil {
		return nil, logs.Wrap(err, "failed to get driver")
	}

	result, err := d.FindMany(request.Database, request.Collection, request.Filter, request.Limit)
	if err != nil {
		return nil, logs.Wrap(err, "failed to find document")
	}
	if result.Error != "" {
		return &result, errors.New(result.Error)
	}
	return &result, nil
}

func ReadOne(request types.OperationRequestReadOne) (*types.OperationResult, error) {
	var requestMany types.OperationRequestReadMany
	requestMany.Database = request.Database
	requestMany.Collection = request.Collection
	requestMany.Filter = request.Filter
	requestMany.Limit = 1
	return ReadMany(requestMany)
}
