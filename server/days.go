package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/timeshibot-backend/store"
	"github.com/rombintu/timeshibot-backend/tools"
)

func (s *Server) UpdateDay() gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chat_id")
		week := c.Param("week")
		day := c.Param("day")

		var subjects []store.Subject

		if err := c.BindJSON(&subjects); err != nil {
			s.respondWithError(c, 401, err)
			return
		}

		if err := s.Database.UpdateDay(chatID, week, day, subjects); err != nil {
			s.respondWithError(c, 500, err)
			return
		}
		c.JSON(http.StatusOK, store.REST{
			Message: tools.Updated, Error: 0},
		)
	}
}
