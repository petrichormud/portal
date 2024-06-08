package handlers

import "net/http"

type FuncWithContext func(c Context) error

type Func func(w http.ResponseWriter, r *http.Request) error

type MiddlewareFunc func(next http.Handler) http.Handler

type MiddlewareFuncWithContext func(next FuncWithContext) FuncWithContext
