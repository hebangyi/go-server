package processConfig

import (
	"fmt"
	"go-server/module"
	"gopkg.in/yaml.v3"
	"os"
)

func initManager(filepath string) {
	// 读取yaml 文件
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取文件错误")
		return
	}

	var config module.Config

	yaml.Unmarshal(data, &config)
	fmt.Println(config.Name)
	fmt.Println(config.Modules[0].Game)
}
