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
			if err := s.Store.CreateGroup(chatID); err != nil {
				c.JSON(http.StatusInternalServerError, store.REST{
					Message: err.Error(), Error: 1},
				)
				return
			}
		} else if action == "get" {
			group, err := s.Store.GetGroup(chatID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, store.REST{
					Message: err.Error(), Error: 1},
				)
				return
			}
			c.JSON(http.StatusOK, group)
			return
		} else if action == "all" {
			groups, err := s.Store.GetGroupAll(chatID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, store.REST{
					Message: err.Error(), Error: 1},
				)
				return
			}
			c.JSON(http.StatusOK, groups)
			return
		} else {
			c.JSON(http.StatusBadGateway, store.REST{
				Message: tools.NotFound, Error: 1},
			)
			return
		}

		c.JSON(http.StatusOK, store.REST{
			Message: tools.Created, Error: 0},
		)
	}
}
