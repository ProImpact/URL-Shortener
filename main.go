package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	"github.com/ProImpact/urlshortener/internal/app"
	"github.com/ProImpact/urlshortener/internal/config"
	"github.com/ProImpact/urlshortener/internal/db"
)

var configFile = flag.String("config", "config.json", "Configuration file")

func main() {
	flag.Parse()
	f, err := os.Open(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fileData, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var cfg config.Configuration
	err = json.Unmarshal(fileData, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	db.MigrateTo(&cfg)
	application := app.NewApplication(&cfg)
	application.Run()
}
