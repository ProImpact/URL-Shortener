package db

import (
	"database/sql"
	"embed"
	"strings"

	"github.com/ProImpact/urlshortener/internal/config"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed sql/schema/*.sql
var embedMigrations embed.FS

func MigrateTo(cfg *config.Configuration) {
	db, err := sql.Open("sqlite", cfg.Database.Path)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}
	if err := goose.UpTo(db, "sql/schema", int64(cfg.Database.Version)); err != nil {
		if strings.Contains(err.Error(), "no migrations to run") {
			return
		}
		panic(err)
	}
}
