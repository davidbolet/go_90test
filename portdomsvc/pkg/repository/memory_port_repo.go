package repository

import (
	"errors"

	"github.com/davidbolet/go_90test/proto"
)

func CreateRepository() Repository {
	return ProtoPortRepository{database: make(map[string]proto.Port)}
}

type ProtoPortRepository struct {
	database map[string]proto.Port
}

// SavePort adds a new Port object to the database
func (r ProtoPortRepository) SavePort(port *proto.Port) error {
	r.database[port.KEY] = *port
	return nil
}

// GetPortByKey returns an existing Port with given key from database
func (r ProtoPortRepository) GetPortByKey(key string) (*proto.Port, error) {
	if val, ok := r.database[key]; ok {
		return &val, nil
	}
	return nil, errors.New("key not found")
}

// DeletePort deletes an existing Port with given key from database
func (r ProtoPortRepository) DeletePort(key string) (*proto.Port, error) {
	if val, ok := r.database[key]; ok {
		delete(r.database, key)
		return &val, nil
	}
	return nil, errors.New("key not found")
}
