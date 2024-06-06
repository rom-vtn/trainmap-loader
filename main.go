package main

import (
	"encoding/json"
	"log"
	"os"

	trainmapdb "github.com/rom-vtn/trainmap-db"
)

// Replace these with your file paths
const DB_PATH = "/path/to/database.db"
const CONFIG_PATH = "/path/to/config.json"

func buildDatabase() error {
	fetcher, err := trainmapdb.NewFetcher(DB_PATH, nil)
	if err != nil {
		return err
	}
	configFileName := CONFIG_PATH
	var config trainmapdb.LoaderConfig
	content, err := os.ReadFile(configFileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		return err
	}
	err = fetcher.LoadDatabase(config)
	if err != nil {
		return err
	}
	log.Default().Println("All feeds have been loaded successfully!")
	return nil
}

func main() {
	err := buildDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
