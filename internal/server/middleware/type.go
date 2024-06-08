package middleware

import (
	"net/http"

	"github.com/petrichormud/portal/internal/server/handler"
)

type Func func(next http.Handler) http.Handler

type FuncWithContext func(next handler.FuncWithContext) handler.FuncWithContext
