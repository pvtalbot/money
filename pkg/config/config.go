package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Database_host string
}

func GetConfig() Configuration {
	jsonFile, err := os.Open("/home/ubuntu/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Configuration

	json.Unmarshal([]byte(byteValue), &config)

	return config
}
