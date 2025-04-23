package model

import (
	"ctf01d/internal/httpserver"
)

func NewSessionFromModel(u *User) *httpserver.SessionResponse {
	return &httpserver.SessionResponse{
		Id:   &u.Id,
		Name: &u.Username,
		Role: (*string)(&u.Role),
	}
}
