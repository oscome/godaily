package day002

import (
	"fmt"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func read() {
	viper.AddConfigPath(".")        // 还可以在工作目录中查找配置
	viper.SetConfigFile("dev.yaml") // 指定配置文件路径(这一句跟下面两行合起来表达的是一个意思)
	// viper.SetConfigName("dev")      // 配置文件名称(无扩展名)
	// viper.SetConfigType("yaml")     // 如果配置文件的名称中没有扩展名，则需要配置此项
	err := viper.ReadInConfig() // 配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func TestViper(t *testing.T) {
	read()
	t.Log(viper.GetString("name"))
	t.Log(viper.GetString("log.level"))
}

func TestWatch(t *testing.T) {
	read()
	t.Log(viper.GetString("name"))
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		read()
		t.Log(viper.GetString("name"))
		t.Log(viper.GetString("log.level"))
	})
	time.Sleep(100 * time.Second)
}
