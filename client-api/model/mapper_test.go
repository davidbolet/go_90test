package model

import (
	"testing"

	"github.com/davidbolet/go_90test/proto"
	"github.com/stretchr/testify/assert"
)

func TestMapPortToProtoHappyPath(t *testing.T) {
	inPort := Port{
		Key:         "1",
		Name:        "in",
		Coordinates: []float32{1.0, 1.0},
		City:        "Barcelona",
		Country:     "Spain",
		Alias:       []string{},
		Regions:     []string{},
		Province:    "Barcelona",
		Timezone:    "GMT+1",
		Unlocs:      []string{"test"},
		Code:        "1",
	}
	expected := proto.Port{
		KEY:         "1",
		Name:        "in",
		Coordinates: []float32{1.0, 1.0},
		City:        "Barcelona",
		Country:     "Spain",
		Alias:       []string{},
		Regions:     []string{},
		Province:    "Barcelona",
		Timezone:    "GMT+1",
		Unlocs:      []string{"test"},
		Code:        "1",
	}
	result, err := MapToProto(&inPort)
	assert.Nil(t, err)
	assert.Equal(t, &expected, result)
}

func TestMapPortToProtoWrongCoordinates(t *testing.T) {
	inPortErr := Port{
		Key:         "1",
		Name:        "in",
		Coordinates: []float32{1.0},
		City:        "Barcelona",
		Country:     "Spain",
		Alias:       []string{},
		Regions:     []string{},
		Province:    "Barcelona",
		Timezone:    "GMT+1",
		Unlocs:      []string{},
		Code:        "1",
	}
	result, err := MapToProto(&inPortErr)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestMapProtoToPortHappyPath(t *testing.T) {
	inPort := proto.Port{
		KEY:         "1",
		Name:        "in",
		Coordinates: []float32{1.0, 1.0},
		City:        "Barcelona",
		Country:     "Spain",
		Alias:       []string{},
		Regions:     []string{},
		Province:    "Barcelona",
		Timezone:    "GMT+1",
		Unlocs:      []string{"test"},
		Code:        "1",
	}
	expected := Port{
		Key:         "1",
		Name:        "in",
		Coordinates: []float32{1.0, 1.0},
		City:        "Barcelona",
		Country:     "Spain",
		Alias:       []string{},
		Regions:     []string{},
		Province:    "Barcelona",
		Timezone:    "GMT+1",
		Unlocs:      []string{"test"},
		Code:        "1",
	}
	result, err := MapToPort(&inPort)
	assert.Nil(t, err)
	assert.Equal(t, &expected, result)
}

func TestMapProtoToPortWrongCoordinates(t *testing.T) {
	inPortErr := proto.Port{
		KEY:         "1",
		Name:        "in",
		Coordinates: []float32{1.0},
		City:        "Barcelona",
		Country:     "Spain",
		Alias:       []string{},
		Regions:     []string{},
		Province:    "Barcelona",
		Timezone:    "GMT+1",
		Unlocs:      []string{},
		Code:        "1",
	}
	result, err := MapToPort(&inPortErr)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
