package settings

import "github.com/spf13/viper"

type Settings struct {
	LoginDB    string
	PasswordDB string
	Host       string
	Port       string
}

func NewSettings() *Settings {
	viper.AddConfigPath("config")
	viper.SetConfigType("env")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	s := &Settings{}
	s.LoginDB = viper.GetString("MONGO_INITDB_ROOT_USERNAME")
	s.PasswordDB = viper.GetString("MONGO_INITDB_ROOT_PASSWORD")
	s.Host = viper.GetString("HOST")
	s.Port = viper.GetString("PORT")

	return s
}
