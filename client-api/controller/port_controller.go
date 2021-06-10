package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/davidbolet/go_90test/client-api/model"
	"github.com/davidbolet/go_90test/client-api/repository"
	"github.com/gin-gonic/gin"
)

type PortController struct {
	repo repository.Repository
}

//NewPortController returns a new instance of the controller
func NewPortController(repo repository.Repository) *PortController {
	return &PortController{repo: repo}
}

//GetPortById handles the get single port command
func (c *PortController) GetPortByKey(ctx *gin.Context) {
	log.Printf("[START] Get Port by key")
	defer log.Printf("[FINISH] Get Port by key")
	key := ctx.Param("key")
	result, err := c.repo.GetPortByKey(key)
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "key not found") {
			ctx.AbortWithStatus(404)
		} else {
			ctx.AbortWithStatus(500)
		}
		return
	}
	if result == nil {
		ctx.AbortWithStatus(404)
		return
	}
	respBody, err := json.Marshal(result)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(500, err)
		return
	}
	ctx.Status(http.StatusOK)
	ctx.Header("Content-Type", gin.MIMEJSON)

	_, err = ctx.Writer.Write(respBody)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	ctx.Writer.Flush()
}

//GetPortById handles the get single port command
func (c *PortController) SavePort(ctx *gin.Context) {
	var port model.Port
	log.Printf("[START] Save Port")
	defer log.Printf("[FINISH] Save Port")
	err := ctx.BindJSON(&port)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(400, err)
		return
	}
	result, err := c.repo.SavePort(&port)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(500, err)
		return
	}
	respBody, err := json.Marshal(result)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(500, err)
		return
	}
	ctx.Status(http.StatusOK)
	ctx.Header("Content-Type", gin.MIMEJSON)

	_, err = ctx.Writer.Write(respBody)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(500, err)
		return
	}
	ctx.Writer.Flush()
}

//GetPortById handles the get single port command
func (c *PortController) DeletePort(ctx *gin.Context) {
	log.Printf("[START] Delete Port by key")
	defer log.Printf("[FINISH] Delete Port by key")

	key := ctx.Param("key")
	result, err := c.repo.DeletePort(key)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(500, err)
	}
	respBody, err := json.Marshal(result)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(500, err)
		return
	}
	ctx.Status(http.StatusOK)
	ctx.Header("Content-Type", gin.MIMEJSON)

	_, err = ctx.Writer.Write(respBody)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(500, err)
		return
	}
	ctx.Writer.Flush()

}
