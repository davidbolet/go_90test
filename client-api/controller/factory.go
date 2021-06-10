package controller

import (
	"github.com/davidbolet/go_90test/client-api/repository"
	"github.com/gin-gonic/gin"
)

// NewAssetWebService creates a new instance of the rest service
func NewPortWebService(repo repository.Repository) *gin.Engine {
	var r *gin.Engine
	gin.SetMode(gin.ReleaseMode)
	r = gin.New()
	portController := NewPortController(repo)

	v1 := r.Group("/v1")

	v1.GET("/port/:key", portController.GetPortByKey)
	v1.POST("/port", portController.SavePort)
	v1.PUT("/port", portController.SavePort)
	v1.DELETE("/port/:key", portController.DeletePort)
	return r
}
