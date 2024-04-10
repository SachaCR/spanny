package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	MigrationFilesPath   string
	ServicePath          string
	Port                 int
	Env                  string
	ProjectId            string
	InstanceId           string
	DatabaseId           string
	UsingSpannerEmulator bool
}

func LoadConfiguration(env string, configPath string) (Config, error) {

	viper.AddConfigPath(configPath)
	viper.SetConfigName(".spannyrc.json")
	viper.SetConfigType("json")

	viper.SetDefault("migrationFilesPath", "./migrations")
	viper.SetDefault("servicePath", "localhost")
	viper.SetDefault("port", 9010)
	viper.SetDefault("envs.default.project", "default-project")
	viper.SetDefault("envs.default.instance", "default-instance")
	viper.SetDefault("envs.default.database", "default-database")
	viper.SetDefault("envs.default.use-emulator", true)

	var err = viper.ReadInConfig()

	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)

		if ok {
			fmt.Println("CONFIG FILE NOT FOUND: Default values loaded")
		} else {
			panic(fmt.Errorf("CONFIG FILE ERROR: %w", err))
		}
	}

	envConfigString := fmt.Sprintf("envs.%s", env)

	envConfig := viper.GetStringMapString(envConfigString)

	if len(envConfig) == 0 {
		fmt.Printf("ENV NOT FOUND ( %s ): Default values loaded\n", env)
		envConfig = viper.GetStringMapString("envs.default")
	}

	usingSpannerEmulator, err := strconv.ParseBool(envConfig["use-emulator"])

	if err != nil {
		usingSpannerEmulator = false
	}

	return Config{
		MigrationFilesPath:   viper.GetString("migrationFilesPath"),
		ServicePath:          viper.GetString("servicePath"),
		Port:                 viper.GetInt("port"),
		Env:                  env,
		ProjectId:            envConfig["project"],
		InstanceId:           envConfig["instance"],
		DatabaseId:           envConfig["database"],
		UsingSpannerEmulator: usingSpannerEmulator,
	}, nil
}
