package main

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
