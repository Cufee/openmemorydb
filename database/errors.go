package database

import "errors"

var (
	ErrNoDocuments = errors.New("no documents found")

	ErrDatabaseExists   = errors.New("database exists")
	ErrDatabaseNotFound = errors.New("database not found")

	ErrCollectionExists   = errors.New("collection exists")
	ErrCollectionNotFound = errors.New("collection not found")
)
