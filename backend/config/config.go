package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

/*
@author: shg
@since: 2022/12/2
@desc: //TODO
*/

type tomlConfig struct {
	DataSource DataSource
	Qiniu      Qiniu
}

type DataSource struct {
	IP       string
	Port     string
	UserName string
	Password string
}

type Qiniu struct {
	QiniuAccessKey string
	QiniuSecretKey string
}

var Config *tomlConfig

func init() {
	// 程序启动时，执行int方法
	Config = new(tomlConfig)
	_, err := toml.DecodeFile("config/config.toml", &Config)
	if err != nil {
		log.Fatal("decode file failed, filename:config/config.toml")
		panic(err)
	}
}
