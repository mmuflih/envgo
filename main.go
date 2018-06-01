package main

import (
	"fmt"

	"github.com/mmuflih/envgo/conf"
)

var config conf.Config

func init() {
	config = conf.NewConfig()
}

func main() {
	fmt.Println(config.GetString("env"))
}
