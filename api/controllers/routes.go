package controllers

import "lets-go/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/api/data", middlewares.SetMiddlewareJSON(s.Tsv)).Methods("GET")
}
