package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	RefreshToken 	string				`json:"refreshToken"`
	Server 			Server				`json:"server"`
	View 			View 				`json:"view"`
}


type View struct {
	HtmlGlob string 				`json:"htmlGlob"`
	LetfDelim string				`json:"letfDelim"`
	RightDelim string				`json:"rightDelim"`
}

type Server struct {
	Port int 						`json:"port"`
}

var Conf = new(Config)

//初始化
func init()  {
	viper.SetConfigFile("./config/config.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()               // 读取配置信息
	if err != nil {                           // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
}
