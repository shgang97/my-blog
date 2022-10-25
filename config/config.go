package config

/*
@author: shg
@since: 2022/10/13
@desc: //TODO
*/
import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type tomlConfig struct {
	Viewer     Viewer
	System     System
	DataSource DataSource
	Comment    Comment
}

// Viewer SystemConfig 从配置文件读取
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

type System struct {
	AppName        string
	Version        float32
	CurrentDir     string
	CdnUrl         string
	QiniuAccessKey string
	QiniuSecretKey string
	//Valine          bool
	//ValineAppId     string
	//ValineAppKey    string
	//ValineServerURL string
}

type DataSource struct {
	IP       string
	Port     string
	UserName string
	Password string
}

type Comment struct {
	Valine          bool
	ValineAppId     string
	ValineAppKey    string
	ValineServerURL string
}

var Config *tomlConfig

func init() {
	// 程序启动时，执行int方法
	Config = new(tomlConfig)
	Config.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Config.System.CurrentDir = currentDir
	_, err := toml.DecodeFile("config/config.toml", &Config)
	if err != nil {
		log.Fatal("decode file failed, filename:config/config.toml")
		panic(err)
	}
}
