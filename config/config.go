package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ModuleName string
	ServiceName string
	ProjectPath string
}

func ReadCommandLineParameters() {
	pflag.String("module-name", "module-name", "New microservice module name (go mod init format)")
	pflag.String("name", "service-name", "New microservice name (e.g. auth)")
	pflag.String("path", ".", "The path where the new microservice will be saved")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.GetString("name")
}

func ValidateParameters() error {
	moduleName := viper.GetString("module-name")
	if moduleName == "" {
		log.Fatalf("empty module name")
	}

	serviceName := viper.GetString("name")
	if serviceName == "" {
		log.Fatalf("empty name")
	}

	pathName := viper.GetString("path")
	if pathName == "" {
		log.Fatalf("empty path")
	}

	return nil
}

func SetConfigParameters() *Config {
	return &Config{
		ModuleName: viper.GetString("module-name"),
		ServiceName: viper.GetString("name"),
		ProjectPath: viper.GetString("path") + "/" + viper.GetString("name"),
	}
}

func NewConfig() *Config {
	ReadCommandLineParameters()
	if err := ValidateParameters(); err != nil{
		log.Fatalf("not valid command line parametr error: %s", err.Error())
	}

	return SetConfigParameters()
}