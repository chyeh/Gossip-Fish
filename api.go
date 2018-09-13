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
	s.httpServer.router.GET("/search/articles", s.searchArticles)
	s.httpServer.router.GET("/search/comments", s.searchComments)
}

func (s *apiServer) searchArticles(c *gin.Context) {
	c.JSON(http.StatusOK, "search articles")
}

func (s *apiServer) searchComments(c *gin.Context) {
	c.JSON(http.StatusOK, "search comments")
}
