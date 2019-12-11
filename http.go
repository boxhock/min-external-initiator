package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebServer struct {
	router *gin.Engine
	service *Service
}

func RunWebserver(service *Service) {
	ws := WebServer{
		service: service,
	}
	ws.createRouter()

	err := ws.router.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func (ws *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws.router.ServeHTTP(w, r)
}

func (ws *WebServer) createRouter() {
	r := gin.Default()
	r.POST("/jobs", ws.CreateJob)
	r.DELETE("/jobs/:jobid", ws.DeleteJob)
	ws.router = r
}

type CreateSubscriptionReq struct {
	JobID  string `json:"jobId"`
	Type   string `json:"type"`
	Params JobConfig `json:"params"`
}

func (ws WebServer) CreateJob(c *gin.Context) {
	var req CreateSubscriptionReq
	err := c.BindJSON(&req)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusInternalServerError, nil)
		return
	}

	ws.service.SubscribeToJob(req.JobID, req.Params)

	c.JSON(http.StatusCreated, nil)
}

func (ws WebServer) DeleteJob(*gin.Context) {
	// TODO: Implement deleting of jobs/stopping subscriptions
}
