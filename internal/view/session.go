package view

import (
	"ctf01d/internal/httpserver"
	"ctf01d/internal/model"
)

func NewSessionFromModel(u *model.User) *httpserver.SessionResponse {
	return &httpserver.SessionResponse{
		Id:   &u.Id,
		Name: &u.Username,
		Role: (*string)(&u.Role),
	}
}
