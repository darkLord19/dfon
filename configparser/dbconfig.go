package configparser

import "encoding/json"

// DBConfig represents a single config object created from config file
type DBConfig struct {
	Name      string  `json:"name"`
	Host      string  `json:"host,omitempty"`
	Port      uint16  `json:"port,omitempty"`
	Type      string  `json:"type,omitempty"`
	User      string  `json:"user,omitempty"`
	Threshold float32 `json:"threshold,omitempty"`
}

// New returns a new DBConfig object
// in: json formatted string
func New(conf string) *DBConfig {
	var dbc DBConfig
	err := json.Unmarshal([]byte(conf), &dbc)
	if err != nil {
		return nil
	}
	return &dbc
}
