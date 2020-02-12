package conf

import (
	"encoding/json"
	"io/ioutil"
)

var (
	ConfigPath string
	Conf = &Config{}
)
//var (
//	// config
//	confPath string
//	client   *conf.Client
//	// Conf .
//
//)

type Config struct {
	MySQLConfig map[string]string `json:"mysql"`
}

//初始化配置
func Load() error {
	data, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, Conf); err != nil {
		return err
	}
	return nil
}
