package aggregator

import (
	"github.com/spf13/viper"
)

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

func GetElementSliceString(s string) []string {
	vi := viper.New()
	vi.SetConfigFile("element.yml")
	vi.AddConfigPath("./")
	err := vi.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return vi.GetStringSlice(s)
}