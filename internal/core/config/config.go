package coreconfig

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	viper   *viper.Viper
	Profile string
}

type PropertyParserFunc func(interface{}) (interface{}, error)

const (
	PROFILE_ENV   = "ENV_PROFILE"
	LOCAL_PROFILE = "LOCAL"
)

func NewConfig() *Config {
	v := viper.New()
	profile, ok := os.LookupEnv(PROFILE_ENV)
	fmt.Printf("Set profile %s\n", profile)
	if ok == false || profile == LOCAL_PROFILE {
		fmt.Println("Inside Local Profile Configuration")
		err := setupLocalEnv(v)
		if err != nil {
			panic(err)
		}
	}
	return &Config{viper: v, Profile: profile}
}

func setupLocalEnv(v *viper.Viper) error {
	v.AddConfigPath(".")
	v.SetConfigFile(".env")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	v.AutomaticEnv()
	return nil
}

func (config *Config) GetValue(key string, parseFunc PropertyParserFunc) (interface{}, error) {
	config.viper.BindEnv(key)
	val := config.viper.Get(key)
	val, err := parseFunc(val)
	if err != nil {
		return nil, errors.New("Error parsing value")
	}
	return val, nil
}

func (config *Config) GetValuesMapString(keys ...string) map[string]string {
	propertyMap := make(map[string]string)
	for _, key := range keys {
		config.viper.BindEnv(key)
		val := config.viper.GetString(key)
		propertyMap[key] = val
	}
	return propertyMap
}

func StringParser() PropertyParserFunc {
	return func(s interface{}) (interface{}, error) {
		property, ok := s.(string)
		if !ok {
			return nil, errors.New("Error Parsing property")
		}
		return property, nil
	}
}
