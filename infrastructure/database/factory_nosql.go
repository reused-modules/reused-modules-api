package database

import (
	"errors"

	"reused-modules-api/adapter/repository"
)

var (
	errInvalidNoSQLDatabaseInstance = errors.New("invalid nosql db instance")
)

const (
	InstanceMongoDB int = iota
)

func NewDatabaseNoSQLFactory(instance int) (repository.NoSQL, error) {
	switch instance {
	case InstanceMongoDB:
		return NewMongoHandler(newConfigMongoDB())
	default:
		return nil, errInvalidNoSQLDatabaseInstance
	}
}
