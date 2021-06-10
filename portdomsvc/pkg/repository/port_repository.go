package repository

import "github.com/davidbolet/go_90test/proto"

type Repository interface {
	SavePort(port *proto.Port) error
	GetPortByKey(key string) (*proto.Port, error)
	DeletePort(key string) (*proto.Port, error)
}
