package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"picture/global"
)

func InitViper() *viper.Viper {
	fmt.Println("初始化 viper")

	var config string
	// 尝试读取命令行参数 默认为config-default.json
	flag.StringVar(&config, "config", "config.json", "choose config file name.")
	flag.Parse()

	// 判空
	if config == "" {
		panic(fmt.Errorf("配置文件名字不能为空"))
	}

	// 创建viper实例
	v := viper.New()

	// 指定配置文件名称
	v.SetConfigName(config)
	// 指定文件类型
	v.SetConfigType("json")
	//配置搜索路径
	v.AddConfigPath("./")
	v.AddConfigPath("./conf/")
	v.AddConfigPath("./config/")
	// 查找并读取配置文件
	err := v.ReadInConfig()
	// 处理读取配置文件的错误
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到；如果需要可以忽略
			panic(fmt.Errorf("配置文件不存在, 未找到%s \n", config))
		} else {
			// 配置文件被找到，但产生了另外的错误
			panic(fmt.Errorf("配置文件被找到，但产生了另外的错误: %s \n", err))
		}
	}
	// 运行时实时读取配置文件
	v.WatchConfig()
	// 配置文件发生变更之后会调用的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件更改:", e.Name)
		if err := v.Unmarshal(&global.PIC_CONFIG); err != nil {
			panic(fmt.Errorf("配置文件格式化错误: %s \n", err))
		}
	})

	// 将配置文件内容转化为结构体存入全局变量
	if err := v.Unmarshal(&global.PIC_CONFIG); err != nil {
		panic(fmt.Errorf("配置文件格式化错误: %s \n", err))
	}

	//fmt.Println(global.PIC_CONFIG)

	fmt.Println("导入配置文件成功")
	return v
}
