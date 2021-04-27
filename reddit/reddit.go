package reddit

import (
	"github.com/spf13/viper"
)

func RedditHello() string {
	return "Doing stuff from reddit..."
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

