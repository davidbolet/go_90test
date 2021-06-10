package mocks

import (
	"github.com/davidbolet/go_90test/client-api/model"
	"github.com/stretchr/testify/mock"
)

type PortRepositoryMock struct {
	mock.Mock
	ToReturn *model.Port
}

func (p *PortRepositoryMock) SavePort(port *model.Port) (*model.Port, error) {
	return port, nil
}

func (p *PortRepositoryMock) GetPortByKey(key string) (*model.Port, error) {
	return p.ToReturn, nil
}

func (p *PortRepositoryMock) DeletePort(key string) (*model.Port, error) {
	return p.ToReturn, nil
}
