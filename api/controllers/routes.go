package controllers

import "github.com/husniramdani/lets-go/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/api/v1/data", middlewares.SetMiddlewareJSON(s.Tsv)).Methods("GET")
}
