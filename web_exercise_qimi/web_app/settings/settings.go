package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/**
  @author: CodeWater
  @since: 2023/11/11
  @desc:
	另外一种读取配置文件的方式： 创建一些结构体，然后使用Unmarshal方法将配置信息直接反序列化到结构体中
**/

func Init() (err error) {
	//viper.SetConfigFile("config.json") // 指定配置文件
	viper.SetConfigName("config")     // 指定配置文件名称（不需要制定配置文件的扩展名）
	viper.SetConfigType("yaml")       // 指定配置文件类型（专用于从远程etcd获取配置信息时指定配置文件类型）
	viper.AddConfigPath("./web_app/") // 指定查找配置文件的路径（这里使用相对路径）
	err = viper.ReadInConfig()        // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	// 监控配置文件的变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return
}
