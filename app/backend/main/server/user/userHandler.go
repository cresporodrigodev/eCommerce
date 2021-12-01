package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNewUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "creando usuario...")
	}
}
