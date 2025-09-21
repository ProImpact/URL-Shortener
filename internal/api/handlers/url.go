package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/ProImpact/urlshortener/internal/api/utils"
	"github.com/ProImpact/urlshortener/internal/db"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type URLHandler struct {
	queries *db.Queries
	host    string
}

type shortUrlRequest struct {
	URL string `json:"url"`
}

func (s shortUrlRequest) Validate() error {
	return validation.ValidateStruct(&s, validation.Field(&s.URL, validation.Required, is.URL))
}

func NewUrlhandler(database *sql.DB, host string) *URLHandler {
	return &URLHandler{
		queries: db.New(database),
		host:    host,
	}
}

func (u *URLHandler) ShortUrl(w http.ResponseWriter, r *http.Request) {
	rb := new(shortUrlRequest)
	err := utils.ParseJSONBody(r, rb)
	if err != nil {
		utils.SendJSON(403, w, err.Error())
		return
	}
	err = rb.Validate()
	if err != nil {
		utils.SendJSON(403, w, map[string]any{
			"error": err.Error(),
		})
		return
	}
	found := true
	_, err = u.queries.URLGetByOriginal(r.Context(), rb.URL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			found = false
		} else {
			utils.SendJSON(500, w, map[string]any{
				"error": "Internal server error",
			})
			slog.Error(err.Error())
			return
		}
	}
	if found {
		utils.SendJSON(500, w, map[string]any{
			"error": "url already registred",
		})
		return
	}
	shortenURL, err := utils.GenURL(r.Context(), rb.URL, u.queries)
	if err != nil {
		utils.SendJSON(500, w, map[string]any{
			"error": "Internal server error",
		})
		slog.Error(err.Error())
		return
	}
	urlCreated, err := u.queries.URLCreate(r.Context(), db.URLCreateParams{
		Original:  rb.URL,
		Shortened: shortenURL,
		Clicks:    0,
		Created:   time.Now(),
		ID:        uuid.NewString(),
	})
	if err != nil {
		utils.SendJSON(500, w, "Internal server error")
		slog.Error(err.Error(), "error", "error when creating the url shortened")
		return
	}
	err = utils.SendJSON(201, w, map[string]any{
		"url": u.host + shortenURL,
		"id":  urlCreated.ID,
	})
	if err != nil {
		utils.SendJSON(500, w, map[string]any{
			"error": "Internal server error",
		})
		slog.Error(err.Error())
		return
	}
}

func (u *URLHandler) GetURl(w http.ResponseWriter, r *http.Request) {
	urlId := r.PathValue("url_id")
	slog.Info("path id", "id", urlId)
	url, err := u.queries.URLGetShortenCode(r.Context(), urlId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.SendJSON(404, w, map[string]any{
				"error": "url not found",
			})
			return
		}
		utils.SendJSON(500, w, map[string]any{
			"error": "Internal server error",
		})
		slog.Error(err.Error())
		return
	}
	http.Redirect(w, r, url.Original, 301)
	go func() {
		u.queries.URLUpdateClicks(context.Background(), url.ID)
	}()
}

func (u *URLHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	// Ejecutar la consulta con limit y offset
	urls, err := u.queries.URLGetAll(r.Context(), db.URLGetAllParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		utils.SendJSON(500, w, map[string]any{
			"error": "Internal server error",
		})
		slog.Error(err.Error())
		return
	}
	utils.SendJSON(200, w, urls)
}
func (u *URLHandler) DeleteEntry(w http.ResponseWriter, r *http.Request) {
	urlId := r.PathValue("url_id")
	err := u.queries.URLDeleteUrlByID(r.Context(), urlId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.SendJSON(404, w, map[string]any{
				"error": "url not found",
			})
			return
		}
		utils.SendJSON(500, w, map[string]any{
			"error": "Internal server error",
		})
		slog.Error(err.Error())
		return
	}
	utils.SendJSON(200, w, map[string]any{
		"message": "url deleted successfully",
	})
}
func (u *URLHandler) GetURLInfo(w http.ResponseWriter, r *http.Request) {
	urlId := r.PathValue("url_id")
	urlInfo, err := u.queries.URLGetByID(r.Context(), urlId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.SendJSON(404, w, map[string]any{
				"error": "url not found",
			})
			return
		}
		utils.SendJSON(500, w, map[string]any{
			"error": "Internal server error",
		})
		slog.Error(err.Error())
		return
	}
	utils.SendJSON(200, w, urlInfo)
}
