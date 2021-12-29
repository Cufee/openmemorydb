package operations

import (
	"errors"

	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/database/driver"
	"byvko.dev/repo/openmemorydb/types"
)

func CreateMany(request types.OperationRequestCreateMany) (*types.OperationResult, error) {
	d, err := driver.GetDriver()
	if err != nil {
		return nil, logs.Wrap(err, "failed to get driver")
	}

	result, err := d.InsertMany(request.Database, request.Collection, request.Documents)
	if err != nil {
		return nil, logs.Wrap(err, "failed to insert document")
	}
	if result.Error != "" {
		return &result, errors.New(result.Error)
	}
	return &result, nil
}

func CreateOne(request types.OperationRequestCreateOne) (*types.OperationResult, error) {
	var requestMany types.OperationRequestCreateMany
	requestMany.Database = request.Database
	requestMany.Collection = request.Collection
	requestMany.Documents = []types.Document{request.Document}
	return CreateMany(requestMany)
}
