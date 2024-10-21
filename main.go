package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	trainmapdb "github.com/rom-vtn/trainmap-db"
	"gorm.io/driver/sqlite"
)

func buildDatabase(configFileName string) error {
	var config trainmapdb.LoaderConfig
	if configFileName == "" {
		return fmt.Errorf("no DB file name in config given")
	}
	content, err := os.ReadFile(configFileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		return err
	}
	dial, useMutex := sqlite.Open(config.DatabasePath), true
	// dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 TimeZone=Europe/Paris"
	// dial, useMutex := postgres.Open(dsn), false
	fetcher, err := trainmapdb.NewFetcher(dial, useMutex, nil)
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
	if len(os.Args) < 2 {
		log.Fatal("Syntax: ./main <config_file.json>")
	}
	configFileName := os.Args[1]
	fmt.Println("Reading config file: ", configFileName)
	err := buildDatabase(configFileName)
	if err != nil {
		log.Fatal(err)
	}
}
