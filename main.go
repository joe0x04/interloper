// no thing to see here yet
package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var config TomlConfig

/**
 * Reads the config file from the filesystem and
 * unmarshals it into a structure
 */
func LoadConfig(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := toml.Decode(string(f), &config); err != nil {
		log.Fatal(err)
	}

	// we don't want credentials in config file, use environment vars
	config.DB.User = os.Getenv("DBUSER")
	config.DB.Pass = os.Getenv("DBPASS")
}

/**
 *
 */
func main() {
	log.Printf("Starting on %s:%d\n", config.HTTP.IPAddress, config.HTTP.Port)
}

/**
 * Runs automatically before main()
 */
func init() {
	LoadConfig("config.toml")
}
