package service

import (
	"context"

	"github.com/davidbolet/go_90test/portdomsvc/pkg/repository"
	"github.com/davidbolet/go_90test/proto"
)

type PortDomainService struct {
	proto.UnimplementedPortDomainServiceServer
	repo repository.Repository
}

func CreateService(repo repository.Repository) *PortDomainService {
	return &PortDomainService{repo: repo}
}

func (s PortDomainService) GetByKey(ctx context.Context, msg *proto.GetPortMsg) (*proto.Port, error) {
	return s.repo.GetPortByKey(msg.KEY)
}
func (s PortDomainService) Save(ctx context.Context, port *proto.Port) (*proto.Port, error) {
	return port, s.repo.SavePort(port)
}

func (s PortDomainService) Delete(ctx context.Context, msg *proto.DeletePortMsg) (*proto.Port, error) {
	return s.repo.DeletePort(msg.KEY)
}
