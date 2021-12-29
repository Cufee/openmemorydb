package operations

import (
	"errors"

	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/database/driver"
	"byvko.dev/repo/openmemorydb/types"
)

func UpdateMany(request types.OperationRequestUpdateMany) (*types.OperationResult, error) {
	d, err := driver.GetDriver()
	if err != nil {
		return nil, logs.Wrap(err, "failed to get driver")
	}

	result, err := d.UpdateMany(request.Database, request.Collection, request.Updates)
	if err != nil {
		return nil, logs.Wrap(err, "failed to update document")
	}
	if result.Error != "" {
		return &result, errors.New(result.Error)
	}
	return &result, nil
}

func UpdateOne(request types.OperationRequestUpdateOne) (*types.OperationResult, error) {
	var requestMany types.OperationRequestUpdateMany
	requestMany.Database = request.Database
	requestMany.Collection = request.Collection
	requestMany.Updates = []types.UpdateOptions{request.Update}
	return UpdateMany(requestMany)
}
