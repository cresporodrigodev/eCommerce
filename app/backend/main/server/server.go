package server

import (
	"fmt"
	"github.com/cresporodrigodev/eCommerce/app/backend/main/server/handler"
	"github.com/cresporodrigodev/eCommerce/app/backend/main/server/user"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddress string
	engine      *gin.Engine
}

func NewServer(host string, port uint) Server {
	server := Server{
		httpAddress: fmt.Sprintf("%s:%d", host, port),
		engine:      gin.New(),
	}

	server.registerRoutes()
	return server
}

func (s *Server) Run() error {
	log.Println("server running on", s.httpAddress)
	return s.engine.Run(s.httpAddress)
}

func (s *Server) registerRoutes() {
	s.engine.GET("handler/", handler.CheckHandler())
	s.engine.GET("user/", user.CreateNewUserHandler())
}
