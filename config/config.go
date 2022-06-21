package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

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

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

//程序启动时就会执行init
func init() {
	Cfg = new(tomlConfig)
	Cfg.System.AppName = "go-blog"
	Cfg.System.Version = 1.0
	currentDir, err := os.Getwd()
    if err != nil {
        log.Println("get path fialed, err:",err)
        return
    }
	Cfg.System.CurrentDir = currentDir
	_, err = toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
