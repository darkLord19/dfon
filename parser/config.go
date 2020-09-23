package parser

import (
	"database/sql"
	"encoding/json"
	"os"
)

// Config is config
type Config struct {
	DBList []Database `json:"databases"`
}

// Database represents a single config object created from config file
type Database struct {
	Name       string   `json:"name"`
	Host       string   `json:"host,omitempty"`
	Port       uint16   `json:"port,omitempty"`
	Type       string   `json:"type,omitempty"`
	User       string   `json:"user,omitempty"`
	Password   string   `json:"password,omitempty"`
	Threshold  float32  `json:"threshold,omitempty"`
	NotifyList []string `json:"notify_list"`
	Connection *sql.DB
}

// LoadConfig returns DBConfig object for all the databases which
// user wants to monitor
func LoadConfig(fname string) (Config, error) {
	var config Config
	configFile, err := os.Open(fname)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}
