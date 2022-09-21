package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/timeshibot-backend/store"
	"github.com/rombintu/timeshibot-backend/tools"
)

type Server struct {
	Router   *gin.Engine
	Database *store.Database
}

func NewServer() *Server {
	return &Server{
		Router:   gin.Default(),
		Database: store.NewDatabase(),
	}
}

func (s *Server) Start() error {
	if err := s.ConfigureDatabase(); err != nil {
		return err
	}
	s.ConfigureRouter()

	return http.ListenAndServe(
		tools.GetEnvOrDefault("PORT", tools.DefaultServerPort),
		s.Router,
	)
}

func (s *Server) ConfigureRouter() {
	s.Router.GET("/", s.Index())
	s.Router.GET("/:chat_id/:action", s.Group())
	s.Router.POST("/:chat_id/:week/:day", s.UpdateDay())
}

func (s *Server) ConfigureDatabase() error {
	if err := s.Database.Open(); err != nil {
		return err
	}
	if err := s.Database.Driver.AutoMigrate(
		&store.Group{},
		&store.TimeTable{},
		&store.Day{},
		&store.Subject{},
	); err != nil {
		return err
	}

	return nil
}

func (s *Server) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, store.REST{
			Message: tools.ServerIsStarting, Error: 0},
		)
	}
}
