package config

type Configuration struct {
	Port     int      `json:"port"`
	Database Database `json:"database"`
}

type Database struct {
	Version int    `json:"version"`
	Path    string `json:"path"`
}
