package models

type SessionResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
