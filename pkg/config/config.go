package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Database_host string `json:"database_host"`
	Front_host    string `json:"front_host"`
}

func GetDatabaseHost() (string, error) {
	config, err := getConfig()
	if err != nil {
		return "", err
	}

	return config.Database_host, nil
}

func GetFrontHost() (string, error) {
	config, err := getConfig()
	if err != nil {
		return "", err
	}

	return config.Front_host, nil
}

func getConfig() (Configuration, error) {
	jsonFile, err := os.Open(("/home/ubuntu/config.json"))
	if err != nil {
		log.Fatal(err)
		return Configuration{}, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
		return Configuration{}, err
	}

	var config Configuration
	err = json.Unmarshal([]byte(byteValue), &config)
	if err != nil {
		log.Fatal(err)
		return Configuration{}, err
	}

	return config, nil
}
