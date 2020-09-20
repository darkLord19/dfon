package configparser

import (
	"encoding/json"
	"os"
)

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
