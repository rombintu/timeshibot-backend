package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/timeshibot-backend/store"
	"github.com/rombintu/timeshibot-backend/tools"
)

func (s *Server) Subjects() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, store.REST{
			Message: tools.Create, Error: 0},
		)
	}
}
