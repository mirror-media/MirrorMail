package configs

import (
	"github.com/spf13/viper"
	"log"
)

var Config AppConfig

type AppConfig struct {
	MailUser  string `mapstructure:"MailUser"`
	MailPwd   string `mapstructure:"MailPwd"`
	JwtSecret string `mapstructure:"JwtSecret"`
}

func LoadConfig() error {
	// configPath string, configName string
	v := viper.New()
	v.SetConfigType("json")
	v.AddConfigPath("./configs")
	v.SetConfigName("config")

	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}
	log.Println("Using config file:", v.ConfigFileUsed())

	if err := v.Unmarshal(&Config); err != nil {
		log.Fatalf("Error unmarshal config file, %s", err)
		return err
	}
	return nil
}
