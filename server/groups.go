package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/timeshibot-backend/store"
	"github.com/rombintu/timeshibot-backend/tools"
)

func (s *Server) Group() gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chat_id")
		action := c.Param("action")

		if action == "create" {
			if err := s.Database.CreateGroup(chatID); err != nil {
				s.respondWithError(c, 500, err)
				return
			}
		} else {
			s.respondWithError(c, 404, tools.NotFound)
			return
		}

		c.JSON(http.StatusOK, store.REST{
			Message: tools.Created, Error: 0},
		)
	}
}
