package collector

import (
	"github.com/spf13/viper"
)

//阿里云秘钥
type Config struct {
	Accesskey   string
	Accesssecrt string
}

var config Config

func getConfig() Config {
	viper.SetConfigName("config") // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(".")      // 第一个搜索路径
	err := viper.ReadInConfig()   // 读取配置数据
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}
