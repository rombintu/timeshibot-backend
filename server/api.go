package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/timeshibot-backend/store"
	"github.com/rombintu/timeshibot-backend/tools"
)

type Server struct {
	Router *gin.Engine
	Store  *store.Store
}

func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
		Store:  store.NewStore(),
	}
}

func (s *Server) Start() error {
	if err := s.ConfigureStore(); err != nil {
		return err
	}
	s.ConfigureRouter()

	return http.ListenAndServe(
		tools.GetEnvOrDefault("PORT", tools.DefaultPortServer),
		s.Router,
	)
}

func (s *Server) ConfigureRouter() {
	s.Router.GET("/", s.Index())
	s.Router.GET("/:chat_id/:week/:day/:action", s.TimetableGET())
	s.Router.POST("/:chat_id/:week/:day/:action", s.TimetablePOST())
}

func (s *Server) ConfigureStore() error {
	if err := s.Store.Open(); err != nil {
		return err
	}
	if err := s.Store.Driver.AutoMigrate(
		&store.Timetable{},
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
