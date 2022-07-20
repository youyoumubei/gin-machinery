package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type BaseConfig struct {
	ConfigPath string `toml:"config_path"`
	AppPort int64 `toml:"app_port"`
	TaskRedisHost string `toml:"task_redis_host"`
	TaskRedisPort int `toml:"task_redis_port"`
	TaskRedisPassword string `toml:"task_redis_password"`
	TaskRedisDB int `toml:"task_redis_db"`
	TaskRedisPoolSize int `toml:"task_redis_pool_size"`
}

var Cfg = &BaseConfig{}

func InitConfig()  {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path += "/config/settings.toml"
	if _, err = toml.DecodeFile(path, &Cfg); err != nil {
		fmt.Println(err)
	}
}