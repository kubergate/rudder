package config

import "github.com/spf13/viper"

func ParseYAML(filePath string) (*viper.Viper, error) {
	// Initialize Viper
	v := viper.New()

	// Set the configuration file's path and format
	v.SetConfigFile(filePath)
	v.SetConfigType("yaml")

	// Read in the configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

// ParseJSON parses a JSON configuration file and returns a Viper configuration object.
func ParseJSON(filePath string) (*viper.Viper, error) {
	// Initialize Viper
	v := viper.New()

	// Set the configuration file's path and format
	v.SetConfigFile(filePath)
	v.SetConfigType("json")

	// Read in the configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

// ParseTOML parses a TOML configuration file and returns a Viper configuration object.
func ParseTOML(filePath string) (*viper.Viper, error) {
	// Initialize Viper
	v := viper.New()

	// Set the configuration file's path and format
	v.SetConfigFile(filePath)
	v.SetConfigType("toml")

	// Read in the configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}
