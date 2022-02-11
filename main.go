// no thing to see here yet
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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
 * We've intercepted a ctrl-c, close file and database handles
 * and exit gracefully
 */
func Shutdown() {
	log.Println("Shutting down...")
}

/**
 *
 */
func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Shutdown()
		os.Exit(1)
	}()

	log.Printf("Starting on %s:%d\n", config.HTTP.IPAddress, config.HTTP.Port)
}

/**
 * Runs automatically before main()
 */
func init() {
	LoadConfig("config.toml")
}
