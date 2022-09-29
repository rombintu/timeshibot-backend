package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/timeshibot-backend/store"
	"github.com/rombintu/timeshibot-backend/tools"
)

func (s *Server) TimetableGET() gin.HandlerFunc { // TODO
	return func(c *gin.Context) {
		chatID := c.Param("chat_id")
		week := c.Param("week")
		day := c.Param("day")
		action := c.Param("action")

		switch action {
		case tools.Get:
			ttFind := store.Timetable{
				ChatID: chatID,
				Name:   day,
				Week:   week,
			}
			tts, err := s.Store.GetTimeTable(ttFind)
			if err != nil {
				c.JSON(http.StatusInternalServerError, store.REST{
					Message: err.Error(), Error: -1},
				)
				return
			}
			c.JSON(http.StatusOK, store.REST{
				Message: tts, Error: 0,
			})
			return
		// case tools.All: // DEV
		// 	tts, err := s.Store.GetTimeTableAll()
		// 	if err != nil {
		// 		c.JSON(http.StatusInternalServerError, store.REST{
		// 			Message: err.Error(), Error: 1},
		// 		)
		// 		return
		// 	}
		// 	c.JSON(http.StatusOK, tts)
		// 	return
		default:
			c.JSON(http.StatusBadGateway, store.REST{
				Message: tools.NotFound, Error: -1},
			)
			return
		}
	}
}

func (s *Server) TimetablePOST() gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("chat_id")
		week := c.Param("week")
		day := c.Param("day")
		action := c.Param("action")

		var subjects []store.Subject

		if err := c.Bind(&subjects); err != nil {
			c.JSON(http.StatusInternalServerError, store.REST{
				Message: err.Error(), Error: 2},
			)
			return
		}

		tt := store.Timetable{
			ChatID: chatID,
			Name:   day,
			Week:   week,
		}

		switch action {
		case tools.Create:
			if err := s.Store.CreateOrUpdateTimeTable(tt, subjects); err != nil {
				c.JSON(http.StatusInternalServerError, store.REST{
					Message: err.Error(), Error: 1},
				)
				return
			}
		// case tools.Update:
		// 	if err := s.Store.UpdateTimeTable(tt, subjects); err != nil {
		// 		c.JSON(http.StatusInternalServerError, store.REST{
		// 			Message: err.Error(), Error: 1},
		// 		)
		// 		return
		// 	}
		default:
			c.JSON(http.StatusBadGateway, store.REST{
				Message: tools.NotFound, Error: 1},
			)
			return
		}

		c.JSON(http.StatusOK, store.REST{
			Message: tools.Create, Error: 0},
		)
	}
}
