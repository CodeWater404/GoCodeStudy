package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/**
  @author: CodeWater
  @since: 2023/11/10
  @desc: viper的基本使用
	https://www.liwenzhou.com/posts/Go/viper/
	特性：
		设置默认值
		从JSON、TOML、YAML、HCL、envfile和Java properties格式的配置文件读取配置信息
		实时监控和重新读取配置文件（可选）
		从环境变量中读取
		从远程配置系统（etcd或Consul）读取并监控配置变化
		从命令行参数读取配置
		从buffer读取配置
		显式配置值
**/

// Config 嵌套配置反序列化演示
type Config struct {
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host   string `mapstructure:"host"`
	DbName string `mapstructure:"dbname"`
	Port   int    `mapstructure:"port"`
}

func main() {
	viper.SetDefault("fileDir", "./")       //设置默认值
	viper.SetConfigName("config")           //配置文件名字
	viper.SetConfigType("yaml")             //配置文件类型
	viper.AddConfigPath("./operate_viper/") //配置文件路径(调用多次以添加多个搜索路径)
	viper.AddConfigPath("/etc/appname/")    //配置文件路径
	viper.AddConfigPath("$HOME/.appname")   //配置文件路径

	err := viper.ReadInConfig() //读取配置信息
	if err != nil {
		panic(fmt.Errorf("fatal error config file:%s \n", err))
	}

	viper.WatchConfig() //监控配置文件变化
	//配置文件发生变化后，会调用回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("viper.Unmarshal failed , err:%v\n", err)
		return
	}
	fmt.Printf("===>config:%#v\n", c)

	//r := gin.Default()
	//r.GET("/version", func(c *gin.Context) {
	//	c.String(http.StatusOK, viper.GetString("version"))
	//})
	//
	//r.Run()
}
