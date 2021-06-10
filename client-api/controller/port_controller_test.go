package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidbolet/go_90test/client-api/model"
	"github.com/davidbolet/go_90test/client-api/pkg/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type PortControllerTestSuite struct {
	suite.Suite
	writerMock *mocks.ResponseWriterMock
	ctx        *gin.Context
	repo       *mocks.PortRepositoryMock
	controller *PortController
	recorder   *httptest.ResponseRecorder
}

func (lrs *PortControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	lrs.writerMock = &mocks.ResponseWriterMock{}
	lrs.recorder = httptest.NewRecorder()
	lrs.ctx, _ = gin.CreateTestContext(lrs.recorder)
	lrs.repo = &mocks.PortRepositoryMock{ToReturn: getTestPort()}
}

func TestLogicalResourceControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PortControllerTestSuite))
}

func (lrs *PortControllerTestSuite) TestSavePort() {
	lrs.controller = NewPortController(lrs.repo)
	reqBody, err := json.Marshal(getTestPort())
	if err != nil {
		log.Fatalf(err.Error())
	}
	lrs.ctx.Request, err = http.NewRequest("POST", "", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatalf(err.Error())
	}
	lrs.controller.SavePort(lrs.ctx)

	lrs.Assert().Equal(http.StatusOK, lrs.recorder.Result().StatusCode, "The return status should be 200")
}

func (lrs *PortControllerTestSuite) TestGetPort() {
	lrs.controller = NewPortController(lrs.repo)
	reqBody, err := json.Marshal(getTestPort())
	if err != nil {
		log.Fatalf(err.Error())
	}
	params := gin.Params{}
	lrs.ctx.Params = append(params, gin.Param{Key: "key", Value: "1"})
	lrs.ctx.Request, err = http.NewRequest("GET", "", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatalf(err.Error())
	}
	lrs.controller.GetPortByKey(lrs.ctx)

	lrs.Assert().Equal(http.StatusOK, lrs.recorder.Result().StatusCode, "The return status should be 200")
}

func (lrs *PortControllerTestSuite) TestGetPortNotExist() {
	var err error
	lrs.controller = NewPortController(lrs.repo)
	lrs.repo.ToReturn = nil
	params := gin.Params{}
	lrs.ctx.Params = append(params, gin.Param{Key: "key", Value: "0"})
	lrs.ctx.Request, err = http.NewRequest("GET", "", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	lrs.controller.GetPortByKey(lrs.ctx)

	lrs.Assert().Equal(http.StatusNotFound, lrs.recorder.Result().StatusCode, "The return status should be 404")
}

func (lrs *PortControllerTestSuite) TestDeletePort() {
	lrs.controller = NewPortController(lrs.repo)
	lrs.repo.ToReturn = getTestPort()
	reqBody, err := json.Marshal(getTestPort())
	if err != nil {
		log.Fatalf(err.Error())
	}
	params := gin.Params{}
	lrs.ctx.Params = append(params, gin.Param{Key: "key", Value: "1"})
	lrs.ctx.Request, err = http.NewRequest("GET", "", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatalf(err.Error())
	}
	lrs.controller.DeletePort(lrs.ctx)

	lrs.Assert().Equal(http.StatusOK, lrs.recorder.Result().StatusCode, "The return status should be 200")
}

func getTestPort() *model.Port {
	return &model.Port{
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
}
