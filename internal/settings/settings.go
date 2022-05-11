package settings

import "github.com/spf13/viper"

type Settings struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	SSLmode  string
}

func NewSettings() *Settings {
	s := &Settings{}
	viper.AddConfigPath("configs")
	viper.SetConfigName("conf")
	viper.SetConfigType("env")
	viper.ReadInConfig()

	s.Host = viper.GetString("HOST")
	s.Port = viper.GetString("PORT")
	s.Password = viper.GetString("POSTGRES_PASSWORD")
	s.User = viper.GetString("USER")
	s.DbName = viper.GetString("DBNAME")

	return s
}
