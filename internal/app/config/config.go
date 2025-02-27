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
		MySQL struct {
			DSN             string `mapstructure:"dsn"`
			MaxOpenConns    int    `mapstructure:"max_open_conns"`
			MaxIdleConns    int    `mapstructure:"max_idle_conns"`
			ConnMaxLifetime string `mapstructure:"conn_max_lifetime"`
			ConnMaxIdleTime string `mapstructure:"conn_max_idle_time"`
		} `mapstructure:"mysql"`
	} `mapstructure:"database"`

	Security struct {
		JWTSecret       string `mapstructure:"jwt_secret"`
		TokenExpiration string `mapstructure:"token_expiration"`
	} `mapstructure:"security"`
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
