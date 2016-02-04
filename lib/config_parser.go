package lib

import (
	"encoding/json"
	"log"
	"os"
	"snow_api/models"
)

func ParseJsonConfig() models.Configuration {
	file, err := os.Open("/tmp/config.json")
	if err != nil {
		log.Fatal("Could not open config file ", err)
	}
	decoder := json.NewDecoder(file)
	config := models.Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Could not decode json ", err)
	}
	return config
}
