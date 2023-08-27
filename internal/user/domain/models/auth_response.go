package models

import customErr "github.com/eli95/template/pkg/custom-errors"

type AuthResponse struct {
	Token string `json:"token"`
}

func (a *AuthResponse) Validate() error {
	if a.Token == "" {
		return customErr.ErrTokenRequired
	}
	return nil
}
