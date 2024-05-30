package common

import (
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	BotToken string `json:"token"`
	Compress bool   `json:"compress"`

	ApiServerPort     int    `json:"api_server_port"`
	WsProxyServerPort int    `json:"ws_proxy_server_port"`
	LogLevel          string `json:"log_level"`
}

const configPath string = "conf.json"

var defaultConf = &Config{
	BotToken:          "Your kook-go bot token",
	Compress:          true,
	ApiServerPort:     9001,
	WsProxyServerPort: 9000,
	LogLevel:          "INFO",
}

func ReadConfig() *Config {

	configData, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Error("read config file failed. generating default file...")
			data, _ := jsoniter.MarshalIndent(defaultConf, "", "  ")
			confFile, _ := os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
			confFile.Write(data)
			os.Exit(1)
		}
	}
	conf := &Config{}
	err = jsoniter.Unmarshal(configData, conf)
	if err != nil {
		log.Panicf("read config file failed, err: %e", err)
	}
	return conf
}
