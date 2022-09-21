package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
