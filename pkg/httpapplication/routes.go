package httpapplication

import "net/http"

func (s *server) SetupRoutes() {
	s.router.HandleFunc("/coinchange", s.httpHandlerCoinChange()).Methods(http.MethodGet)
}
