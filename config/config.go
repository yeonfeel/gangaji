package config

import (
	"time"

	"github.com/yeonfeel/gopkg/config"
)

// ConfigFileName is a config filename of application
const ConfigFileName string = ".gangaji.yaml"

// Config struct contains configuration
type Config struct {
	Server Server
}

// Server struct contains a server configuration
type Server struct {
	GRPCPort                int
	HTTPPort                int
	SSL                     bool
	SSLCertPath             string
	SSLKeyPath              string
	GracefulShutdownTimeout time.Duration
}

// Data is a data of conifg
var data *config.Config

// Default is default values of conifg
var Default = Config{
	Server: Server{
		GRPCPort:                9800,
		HTTPPort:                9801,
		SSL:                     false,
		GracefulShutdownTimeout: (time.Second * 5),
	},
}

func checkMissingResourceEnvvars(data interface{}) {
	// conf := data.(*Config)
	// envutil.SetEnvB(&conf.Resource.Deployment, "RESOURCE_DEPLOYMENT")
}

func init() {
	conf := Default
	data = config.New(ConfigFileName, &conf, checkMissingResourceEnvvars)
}

// Load a config
func Load() *Config {
	data.Load()
	return data.Value.(*Config)
}

// Write a config
func Write() error {
	return data.Write()
}
