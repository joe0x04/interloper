package main

/**
 * The main config file is parsed into this struct
 */
type TomlConfig struct {
	HTTP struct {
		Enabled   bool   `toml:"enabled"`
		IPAddress string `toml:"ip"`
		Port      int    `toml:"port"`
	} `toml:"http"`

	DB struct {
		Host   string `toml:"host"`
		User   string `toml:"username"`
		Pass   string `toml:"password"`
		Schema string `toml:"schema"`
	} `toml:"database"`
}
