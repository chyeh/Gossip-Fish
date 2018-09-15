package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type apiServerConfig struct {
	host          string
	port          string
	elasticsearch string
}

func (c *apiServerConfig) load(cfg *applicationConfig) {
	c.host = cfg.host
	c.port = cfg.port
	c.elasticsearch = cfg.elasticsearch
}

type apiServer struct {
	httpServer *ginHTTPServer
	address    string
	datastore  datastore
}

func newAPIServer(cfg apiServerConfig) *apiServer {
	httpServer := newGinHTTPServer()
	apiServer := &apiServer{
		httpServer: httpServer,
		address:    net.JoinHostPort(cfg.host, cfg.port),
		datastore:  newElasticsearchClient(cfg.elasticsearch),
	}
	apiServer.routes()
	return apiServer
}

func (s *apiServer) run() {
	fmt.Println("starting http service on", s.address)
	s.httpServer.run(s.address)
}

func (s *apiServer) shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		panic(err)
	}
	fmt.Println("service stopped")
}

func (s *apiServer) routes() {
	s.httpServer.router.GET("/search/articles", s.getSearchArticles)
	s.httpServer.router.GET("/search/comments", s.getSearchComments)
}

func (s *apiServer) getSearchArticles(c *gin.Context) {
	query := newQuery()
	if err := c.ShouldBindQuery(query); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := validate.Struct(query); err != nil {
		panic(err)
	}
	result := s.datastore.searchArticles(query)
	c.JSON(http.StatusOK, result)
}

func (s *apiServer) getSearchComments(c *gin.Context) {
	query := newQuery()
	if err := c.ShouldBindQuery(query); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := validate.Struct(query); err != nil {
		panic(err)
	}
	result := s.datastore.searchComments(query)
	c.JSON(http.StatusOK, result)
}
