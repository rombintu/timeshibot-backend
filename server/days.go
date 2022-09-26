package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/timeshibot-backend/store"
	"github.com/rombintu/timeshibot-backend/tools"
)

func (s *Server) UpdateDay() gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chat_id")
		week, err := strconv.Atoi(c.Param("week"))
		if err != nil {
			c.JSON(http.StatusBadRequest, store.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}
		day, err := strconv.Atoi(c.Param("day"))
		if err != nil {
			c.JSON(http.StatusBadRequest, store.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}
		var subjects []store.Subject

		if err := c.BindJSON(&subjects); err != nil {
			c.JSON(http.StatusBadRequest, store.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}

		if err := s.Store.UpdateDay(chatID, week, day, subjects); err != nil {
			c.JSON(http.StatusInternalServerError, store.REST{
				Message: err.Error(), Error: 1},
			)
			return
		}
		c.JSON(http.StatusOK, store.REST{
			Message: tools.Updated, Error: 0},
		)
	}
}
