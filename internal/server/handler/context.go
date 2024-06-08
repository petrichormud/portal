package handler

import "net/http"

type ContextPlayer struct {
	ID int64
	// TODO: Permissions
}

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Player  ContextPlayer
}
