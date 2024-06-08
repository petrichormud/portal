package handler

import "net/http"

type FuncWithContext func(c Context) error

type Func func(w http.ResponseWriter, r *http.Request) error
