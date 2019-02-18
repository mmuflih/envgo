package conf

/*
 * Created by M. Muflih Kholidin
 * Thu May 03 2018 14:06:46
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config interface {
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetStrings(key string) []string
	GetStringSlice(key string) []string
	Init()
}

type viperConfig struct{}

func (v *viperConfig) Init() {
	viper.SetEnvPrefix(`go-clean`)
	viper.AutomaticEnv()

	osEnv := os.Getenv("OS_ENV")
	env := "env"
	if osEnv != "" {
		env = osEnv
	}

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(env + `.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v *viperConfig) GetStringSlice(key string) (c []string) {
	c = viper.GetStringSlice(key)
	return
}

func (v *viperConfig) GetStrings(key string) (c []string) {
	val := viper.GetString(key)
	c = strings.Split(val, ",")
	return
}

func NewConfig() Config {
	v := &viperConfig{}
	v.Init()
	return v
}
