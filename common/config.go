package common

import (
	"time"

	"github.com/BurntSushi/toml"

	e "wangqingang/cunxun/error"
)

// Configs 全局配置信息
// Duration 类型支持的单位h-小时，m-分钟，s-秒
type Configs struct {
	ReleaseMode    bool
	Listen         string
	Gomaxprocs     int
	CoroutineCount int
}

// Config 全局配置信息
var Config *Configs

// InitConfig 加载配置
func InitConfig(path string) error {
	config, err := loadConfig(path)
	if err != nil {
		return err
	}
	Config = config
	return nil
}

func loadConfig(path string) (*Configs, error) {
	config := new(Configs)
	if _, err := toml.DecodeFile(path, config); err != nil {
		return nil, e.SP(e.MConfigErr, e.ConfigParseErr, err)
	}

	return config, nil
}
