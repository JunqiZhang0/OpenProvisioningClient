package config

// Config defines the structure of configuration
type Config struct {
	// Page defines the content of web page
	Page struct {
		Title   string `yaml:"title"`
		Content string `yaml:"content"`
	} `yaml:page`
	// Server defines the server hosting the website
	Server struct {
		// Provider of server, could be "ppe", "aws", etc.
		Provider string `yaml:"provider"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:server`
}

// NewConfig returns a pointer pointing to a new empty config
func NewConfig() *Config {
	return &Config{}
}
