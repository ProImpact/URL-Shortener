package app

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ProImpact/urlshortener/internal/api"
	"github.com/ProImpact/urlshortener/internal/config"
)

type Application struct {
	srv          *http.Server
	router       *http.ServeMux
	cleanupFuncs []func() error
}

func NewApplication(cfg *config.Configuration) *Application {
	a := new(Application)
	a.router = http.NewServeMux()
	db, err := sql.Open("sqlite", cfg.Database.Path)
	if err != nil {
		panic(err)
	}
	a.cleanupFuncs = append(a.cleanupFuncs, db.Close)
	router := api.NewRouter(db, fmt.Sprintf("http://localhost:%d/", cfg.Port))
	router.SetRoutes(a.router)
	a.srv = &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        a.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return a
}

func (a *Application) Run() {
	singTerm := make(chan os.Signal, 1)
	signal.Notify(singTerm, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{})
	go func() {
		slog.Info("app running", "port", a.srv.Addr)
		if err := a.srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		<-singTerm
		fmt.Println("App shutdown")
		err := a.srv.Close()
		if err != nil {
			slog.Error(err.Error())
		}
		for _, clean := range a.cleanupFuncs {
			err = clean()
			if err != nil {
				slog.Error(err.Error())
			}
		}
		done <- struct{}{}
	}()
	<-done
}
