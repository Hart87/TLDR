package twitter

import (
	"github.com/spf13/viper"
)


func TwitterHello() string {
	return "Doing stuff from twitter..."
}

func GetElementString(s string) string {
	vi := viper.New()
	vi.SetConfigFile("element.yml")
	vi.ReadInConfig()
	return vi.GetString(s)
}

func GetElementInt(s string) int {
	vi := viper.New()
	vi.SetConfigFile("element.yml")
	vi.ReadInConfig()
	return vi.GetInt(s)
}