package api

import (
	"api-gateway/config"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	return gin.Default()
}
