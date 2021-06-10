package repository

import "github.com/davidbolet/go_90test/client-api/model"

type Repository interface {
	SavePort(port *model.Port) (*model.Port, error)
	GetPortByKey(key string) (*model.Port, error)
	DeletePort(key string) (*model.Port, error)
}
