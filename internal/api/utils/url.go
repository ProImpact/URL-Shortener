package utils

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"errors"

	"github.com/ProImpact/urlshortener/internal/db"
)

// genURL Handles collition detection when generating shorten urls
func GenURL(ctx context.Context, urlName string, q *db.Queries) (string, error) {
	for {
		urlName += rand.Text()
		shorten := sha512.Sum512([]byte(urlName))
		encoded := hex.EncodeToString(shorten[:])
		_, err := q.URLGetShortenCode(ctx, encoded[:10])
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return encoded[:10], nil
			}
			return "", err
		}
	}
}
