// no thing to see here yet
package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var config TomlConfig

func LoadConfig(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	if _, err := toml.Decode(string(f), &config); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Printf("Starting on %s:%d\n", config.HTTP.IPAddress, config.HTTP.Port)
}

func init() {
	LoadConfig("config.toml")
}
