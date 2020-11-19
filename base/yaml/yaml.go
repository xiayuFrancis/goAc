package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func main1() {
	config, err := yaml.ReadFile("./yaml/conf.yaml")
	if err != nil {
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}
