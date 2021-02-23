package conf

/*
 * Created by M. Muflih Kholidin
 * Thu May 03 2018 14:06:46
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config interface {
	SetConfig(key string, value interface{})
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetStrings(key string) []string
	GetStringSlice(key string) []string
	Init(string)
}

type viperConfig struct{}

func (v *viperConfig) Init(prefix string) {
	viper.SetEnvPrefix(`go-clean`)
	viper.AutomaticEnv()

	osEnv := os.Getenv("OS_ENV")
	env := "env"
	if osEnv != "" {
		env = osEnv
	}

	if prefix != "" {
		env = prefix + "." + env
	}
	envPath := v.getPath()
	fmt.Println("Environment path:", envPath)
	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(envPath + env + `.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func (v *viperConfig) SetConfig(key string, value interface{}) {
	viper.Set(key, value)
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
	v.Init("")
	return v
}

func NewWithPrefix() Config {
	v := &viperConfig{}
	prefix := v.GetString("prefix")
	fmt.Println("Init config with prefix", prefix)
	v.Init(prefix)
	return v
}

func (v *viperConfig) getPath() string {
	args := os.Args
	fmt.Println(args)
	for _, arg := range args {
		if strings.Contains(arg, "--env") {
			envs := strings.Split(arg, "=")
			if len(envs) > 1 {
				path := envs[1]
				if path[len(path)-1:] == "/" {
					return path
				} else {
					return path + "/"
				}
			}
		}
	}
	return ""
}
