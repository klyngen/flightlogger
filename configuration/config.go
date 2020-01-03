package configuration

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Add more good places to hide configs right here ;)
// The order is important. Look exceedingly more general areas
var configpaths = [...]string{".", "$HOME/.flightlogger", "/etc/flightlogger"}
var homeSigns = [...]string{"~", "$HOME"}

// ApplicationConfig contains the config parameters used for this excellent app
type ApplicationConfig struct {
	Serverport            string
	ServerURL             string
	PrivateKeyPath        string
	PublicKeyPath         string
	Tokenexpiration       int
	DatabaseConfiguration DatabaseConfig     `mapstructure:"database"`
	RedisConfiguration    DatabaseConfig     `mapstructure:"redis"`
	EmailConfiguration    EmailConfiguration `mapstructure:"email"`
}

// EmailConfiguration is just that. Configures SMTP-interaction
type EmailConfiguration struct {
	SmtpServer string
	Username   string
	Password   string
	Port       string
}

// DatabaseConfig describes parameters used to connect to a database
type DatabaseConfig struct {
	Hostname string
	Password string
	Database string
	Port     string
	Username string
}

// IsEmpty returns false if no parameters are set
func (d *DatabaseConfig) IsEmpty() bool {
	if len(d.Hostname) > 0 || len(d.Password) > 0 || len(d.Database) > 0 || len(d.Port) > 0 || len(d.Username) > 0 {
		return false
	}

	return true
}

// GetConfiguration - well... it gets the configuration
func GetConfiguration() ApplicationConfig {
	initializeConfig()

	var c ApplicationConfig

	viper.Unmarshal(&c)

	c.PrivateKeyPath = getAbsoulePath(c.PrivateKeyPath)
	c.PublicKeyPath = getAbsoulePath(c.PublicKeyPath)

	return c
}

func getAbsoulePath(path string) string {
	// We dont mess with correct paths
	if filepath.IsAbs(path) {
		return path
	}

	// We have a ~ or $HOME etc
	homeSymbol := hasHomeSymbol(path)
	if len(homeSymbol) > 0 {
		homePath, err := os.UserHomeDir()

		if err != nil {
			log.Fatal("Unable to get the configuration path. We need more permissions")
		}

		// remove the symbol and use the user-home directory
		return filepath.Join(homePath, strings.TrimLeft(path, homeSymbol))
	}

	absolute, err := filepath.Abs(path)

	if err == nil {
		return absolute
	}

	return path
}

func hasHomeSymbol(path string) string {
	for _, s := range homeSigns {
		if strings.HasPrefix(path, s) {
			return s
		}
	}
	return ""
}

func initializeConfig() {
	// what is the name of the config-file
	// Can be any of many configuration parameters
	viper.SetConfigName("flightlog")

	// Where can configs be located
	for _, path := range configpaths {
		viper.AddConfigPath(path)
	}

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Unable to find flightlog-config in any of %v \n And got the following error: \n %v", configpaths, err)
	}
}
