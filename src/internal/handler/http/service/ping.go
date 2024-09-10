package service

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Result string `json:"result"`
}

func (s *Service) Ping(w http.ResponseWriter, _ *http.Request) {
	response := PingResponse{"pong"}
	responseJson, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(500)
	}

	_, err = w.Write(responseJson)
	if err != nil {
		return
	}

	w.WriteHeader(200)
}
