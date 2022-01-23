package core

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	fmt.Println("初始化 viper")

	// 优先级: 命令行 > 环境变量 > config文件
	var config string
	// 尝试读取命令行参数
	flag.StringVar(&config, "config", "./config.json", "choose config file.")
	flag.Parse()

	fmt.Printf("%v", config)

	// 指定配置文件路径
	//viper.SetConfigFile("./config.json")
	//// 查找并读取配置文件
	//err := viper.ReadInConfig()
	//// 处理读取配置文件的错误
	//if err != nil {
	//	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	//		// 配置文件未找到；如果需要可以忽略
	//		panic(fmt.Errorf("配置文件不存在, 未找到config.json \n"))
	//	} else {
	//		// 配置文件被找到，但产生了另外的错误
	//		panic(fmt.Errorf("配置文件被找到，但产生了另外的错误: %s \n", err))
	//	}
	//}
	//
	v := viper.New()
	//v.WatchConfig()
	return v
}
