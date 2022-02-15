// no thing to see here yet
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

var config TomlConfig
var database *sql.DB

/**
 * Reads the config file from the filesystem and
 * unmarshals it into a structure
 */
func LoadConfig(filename string) {
	log.Printf("Loading config from: %s\n", filename)
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
	// catch signals
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Shutdown()
		os.Exit(1)
	}()

	log.Printf("Starting on %s:%d\n", config.HTTP.IPAddress, config.HTTP.Port)
	addr := fmt.Sprintf("%s:%d", config.HTTP.IPAddress, config.HTTP.Port)

	r := mux.NewRouter()
	LoadRoutes(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	srv.ListenAndServe()
}

/**
 * Runs automatically before main()
 */
func init() {
	config := flag.String("config", "config.toml", "The config file for ports and whatnot")
	flag.Parse()

	LoadConfig(*config)
	DBConnect()
}
