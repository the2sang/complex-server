package middleware

import (
  "net/http"

  "github.com/the2sang/complex-server/config"
)

func RegisterMiddleware(mux *http.ServeMux, c config.AppConfig) http.Handler {
  return loggingMiddleware(panicMiddleware(mux, c), c)
}
