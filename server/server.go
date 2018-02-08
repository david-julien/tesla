package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/davidjulien/tesla/config"
	"github.com/davidjulien/tesla/data"
	"github.com/gorilla/mux"
)

// Server represents the HTTP server that provides a REST API
// interface to the database.
type Server struct {
	Router *mux.Router
	Server *http.Server
	addr   string
	Dal    *data.DAL
	log    *log.Entry
}

// New returns a new instance of the HTTP server based on a config.
func New(c *config.Config, dal *data.DAL, entry *log.Entry) *Server {
	router := mux.NewRouter().StrictSlash(true)
	addr := c.Host + ":" + c.Port
	server := &http.Server{
		Addr: addr,
	}

	s := &Server{
		Router: router,
		Server: server,
		addr:   addr,
		Dal:    dal,
		log:    entry,
	}

	return s
}

// Start starts up the server.
func (s *Server) Start() error {
	s.log.Info("Starting server on ", s.addr)
	err := http.ListenAndServe(s.addr, s.Router)
	return err
}
