package main

import (
	"fmt"
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"time"
	"github.com/johnmcdnl/auth/auth"
)

const port int = 8600

func main() {
	authServer()
}

func authServer() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CloseNotify)
	r.Use(middleware.Timeout(60 * time.Minute))

	r.Mount("/debug", middleware.Profiler())

	r.Mount("/auth", auth.AuthRouter())

	http.ListenAndServe(fmt.Sprint(":", port), r)
}
