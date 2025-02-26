package config

import (
	// "log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Env string `mapstructure:"env"`
	} `mapstructure:"app"`
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"database"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)     // 配置文件路径 (如 ./configs)
	viper.SetConfigName("config") // 文件名 (自动识别拓展名)
	viper.SetConfigType("yaml")   // 明确指定类型 (可选)

	// 读取环境变量 (优先级高于配置文件)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
