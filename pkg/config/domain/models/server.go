package models

import customErr "github.com/eli95/template/pkg/custom-errors"

type Server struct {
	Port string `json:"port"`
}

func (s *Server) Validate() error {
	if s.Port == "" {
		return customErr.ErrServerPortRequired
	}
	return nil
}
