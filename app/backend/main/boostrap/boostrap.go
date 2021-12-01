package boostrap

import (
	"github.com/cresporodrigodev/eCommerce/app/backend/main/server"
)

const (
	localHost = "localhost"
	port      = 8080
)

func RunServer() error {
	srv := server.NewServer(localHost, port)
	return srv.Run()
}
