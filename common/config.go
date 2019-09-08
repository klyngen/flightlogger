package common

// ApplicationConfig contains the config parameters used for this excellent app
type ApplicationConfig struct {
	Serverport            string
	DatabaseConfiguration DatabaseConfig `mapstructure:"database"`
}

// DatabaseConfig describes parameters used to connect to a database
type DatabaseConfig struct {
	Hostname string
	Password string
	Database string
	Port     string
	Username string
}
