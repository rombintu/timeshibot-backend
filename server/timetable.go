package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/timeshibot-backend/store"
	"github.com/rombintu/timeshibot-backend/tools"
)

func (s *Server) Timetable() gin.HandlerFunc {
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

// func (s *Server) UpdateDay() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		chatID := c.Param("chat_id")
// 		week, err := strconv.Atoi(c.Param("week"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, store.REST{
// 				Message: err.Error(), Error: 1},
// 			)
// 			return
// 		}
// 		day, err := strconv.Atoi(c.Param("day"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, store.REST{
// 				Message: err.Error(), Error: 1},
// 			)
// 			return
// 		}
// 		var subjects []store.Subject

// 		if err := c.BindJSON(&subjects); err != nil {
// 			c.JSON(http.StatusBadRequest, store.REST{
// 				Message: err.Error(), Error: 1},
// 			)
// 			return
// 		}

// 		// if err := s.Store.UpdateDay(chatID, week, day, subjects); err != nil {
// 		// 	c.JSON(http.StatusInternalServerError, store.REST{
// 		// 		Message: err.Error(), Error: 1},
// 		// 	)
// 		// 	return
// 		// }
// 		c.JSON(http.StatusOK, store.REST{
// 			Message: tools.Updated, Error: 0},
// 		)
// 	}
// }
