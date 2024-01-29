package main

import (
	"fmt"
	"log"

	"github.com/pelletier/go-toml"
)

// Config structure to hold the configuration values
type Config struct {
	Title    string
	Author   string
	Debug    bool
	Database DatabaseConfig
}

// DatabaseConfig structure to hold database configuration values
type DatabaseConfig struct {
	Server string
	Port   int
	DBName string
}

func main() {
	// Read and parse the TOML file
	config, err := readConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	// Print the configuration values
	fmt.Printf("Title: %s\n", config.Title)
	fmt.Printf("Author: %s\n", config.Author)
	fmt.Printf("Debug: %v\n", config.Debug)
	fmt.Printf("Database Server: %s\n", config.Database.Server)
	fmt.Printf("Database Port: %d\n", config.Database.Port)
	fmt.Printf("Database Name: %s\n", config.Database.DBName)
}

func readConfig(filePath string) (*Config, error) {
	// Create a new Config struct
	config := &Config{}

	// Open and parse the TOML file
	tree, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the TOML data into the Config struct
	if err := tree.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
