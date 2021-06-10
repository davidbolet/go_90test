package repository

import (
	"context"
	"log"

	"github.com/davidbolet/go_90test/client-api/model"
	"github.com/davidbolet/go_90test/proto"
	"google.golang.org/grpc"
)

type PortRepository struct {
	client proto.PortDomainServiceClient
	ctx    context.Context
}

func NewPortRepository(addr string) *PortRepository {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := proto.NewPortDomainServiceClient(conn)
	return &PortRepository{client: client, ctx: context.Background()}
}

func (p *PortRepository) SavePort(port *model.Port) (*model.Port, error) {
	outPort, err := model.MapToProto(port)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := p.client.Save(p.ctx, outPort)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result, err := model.MapToPort(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (p *PortRepository) GetPortByKey(key string) (*model.Port, error) {
	res, err := p.client.GetByKey(p.ctx, &proto.GetPortMsg{KEY: key})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result, err := model.MapToPort(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (p *PortRepository) DeletePort(key string) (*model.Port, error) {
	removed, err := p.client.Delete(p.ctx, &proto.DeletePortMsg{KEY: key})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result, err := model.MapToPort(removed)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
