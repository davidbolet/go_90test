package model

import (
	"errors"

	"github.com/davidbolet/go_90test/proto"
)

func MapToProto(inPort *Port) (*proto.Port, error) {
	result := proto.Port{}
	result.KEY = inPort.Key
	result.Name = inPort.Name
	result.Coordinates = make([]float32, 2)
	if len(inPort.Coordinates) != 0 && len(inPort.Coordinates) != 2 {
		return nil, errors.New("invalid coordinates size")
	}
	if len(inPort.Coordinates) == 2 {
		result.Coordinates[0] = inPort.Coordinates[0]
		result.Coordinates[1] = inPort.Coordinates[1]
	}
	result.City = inPort.City
	result.Province = inPort.Province
	result.Alias = []string{}
	result.Regions = []string{}
	result.Timezone = inPort.Timezone
	result.Unlocs = []string{}
	result.Code = inPort.Code
	result.Country = inPort.Country
	for _, v := range inPort.Unlocs {
		result.Unlocs = append(result.Unlocs, v)
	}
	return &result, nil
}

func MapToPort(prt *proto.Port) (*Port, error) {
	result := Port{}
	result.Key = prt.KEY
	result.Name = prt.Name
	result.Coordinates = make([]float32, 2)
	if len(prt.Coordinates) < 2 {
		return nil, errors.New("invalid coordinates size")
	}
	result.Coordinates[0] = prt.Coordinates[0]
	result.Coordinates[1] = prt.Coordinates[1]
	result.City = prt.City
	result.Alias = []string{}
	result.Regions = []string{}
	result.Timezone = prt.Timezone
	result.Unlocs = []string{}
	result.Country = prt.Country
	result.Province = prt.Province
	result.Code = prt.Code

	for _, v := range prt.Unlocs {
		result.Unlocs = append(result.Unlocs, v)
	}
	return &result, nil
}
