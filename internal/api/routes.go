package api

import (
	"database/sql"
	"net/http"

	"github.com/ProImpact/urlshortener/internal/api/handlers"
	"github.com/ProImpact/urlshortener/internal/api/utils"
)

type Router struct {
	urlHanlder *handlers.URLHandler
}

func NewRouter(db *sql.DB, host string) *Router {
	return &Router{
		urlHanlder: handlers.NewUrlhandler(db, host),
	}
}

func (r *Router) SetRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		utils.SendJSON(200, w, map[string]any{
			"status": "ok",
		})
	})
	router.HandleFunc("POST /short", r.urlHanlder.ShortUrl)
	router.HandleFunc("GET /{url_id}", r.urlHanlder.GetURl)
	router.HandleFunc("GET /api/urls/{url_id}", r.urlHanlder.GetURLInfo)
	router.HandleFunc("GET /api/urls", r.urlHanlder.GetAll)
	router.HandleFunc("DELETE /api/urls/{url_id}", r.urlHanlder.DeleteEntry)
}
