package main

//
// Holds data that makes up a community
//
type Community struct {
	id             int
	uuid           string
	fullname       string
	shortname      string
	date_created   int
	date_lastvisit int
	post_count     int
	feature_mask   int
	creator        int
}

//
// The main config file is parsed into this struct
//
type TomlConfig struct {
	HTTP struct {
		Enabled   bool   `toml:"enabled"`
		IPAddress string `toml:"ip"`
		Port      int    `toml:"port"`
		Logfile   string `toml:"logfile"`
	} `toml:"http"`

	DB struct {
		Host   string `toml:"host"`
		User   string `toml:"username"`
		Pass   string `toml:"password"`
		Schema string `toml:"schema"`
	} `toml:"database"`
}
