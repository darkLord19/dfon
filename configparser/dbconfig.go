package configparser

import (
	"encoding/json"
)

// Config is config
type Config struct {
	DBList []Database `json:"databases"`
}

// Database represents a single config object created from config file
type Database struct {
	Name      string  `json:"name"`
	Host      string  `json:"host,omitempty"`
	Port      uint16  `json:"port,omitempty"`
	Type      string  `json:"type,omitempty"`
	User      string  `json:"user,omitempty"`
	Threshold float32 `json:"threshold,omitempty"`
}

// New returns a new Database object
// in: json formatted string
func New(conf string) *Database {
	var dbc Database
	err := json.Unmarshal([]byte(conf), &dbc)
	if err != nil {
		return nil
	}
	return &dbc
}
