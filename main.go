package main

import (
	"fmt"
	"go-server/module"
	yaml "gopkg.in/yaml.v3"
	"os"
)

func main() {
	// 读取yaml 文件
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("读取文件错误")
		return
	}

	var config module.Config
	yaml.Unmarshal(data, &config)
	fmt.Println(config.Name)
	fmt.Println(config.Modules)
}
