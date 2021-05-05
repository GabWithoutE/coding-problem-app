package http

import (
	"encoding/json"
	"github.com/gabriellukechen/coding-problem-app/pkg/helpers"
	"github.com/gabriellukechen/coding-problem-app/pkg/solving"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (s *server) respond(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)

	if data != nil {
		response, _ := json.Marshal(data)
		_, _ = w.Write(response)
	}
}

func (s *server) handleErrorEvents(w http.ResponseWriter, error solving.SolverError) {
}

func (s *server) httpHandlerCoinChange() http.HandlerFunc {
	log := s.Logger()

	return func(w http.ResponseWriter, r *http.Request) {
		d, _ := r.URL.Query()["denominations"]

		di, err := helpers.Atoiarray(d)
		if err != nil {
			log.Info("httpHandlerCoinChange: invalid denominations in query params", zap.Strings("denominations", d))
			s.respond(w, nil, http.StatusBadRequest)
			return
		}

		t, _ := r.URL.Query()["total"]
		if len(t) < 1 {
			log.Info("httpHandlerCoinChange: no total given in query params")
			s.respond(w, nil, http.StatusBadRequest)
			return
		}

		ti, err := strconv.Atoi(t[0])
		if err != nil {
			log.Info("httpHandlerCoinChange: invalid total in query params", zap.Strings("total", t))
			s.respond(w, nil, http.StatusBadRequest)
			return
		}

		p := solving.NewCoinChangeProblem(di, ti)

		solution, err := p.Solve()
		if err != nil {
			log.Info("httpHandlerCoinChange: solve failed", zap.Error(err))
			s.respond(w, nil, http.StatusBadRequest)
			return
		}

		s.respond(w, solution, http.StatusOK)
	}
}
