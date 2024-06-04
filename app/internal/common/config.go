package common

import (
	_ "embed"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	BotToken string `json:"token"`
	Compress bool   `json:"compress"`

	ApiServerPort            int    `json:"api_server_port"`
	WsProxyServerPort        int    `json:"ws_proxy_server_port"`
	LogLevel                 string `json:"log_level"`
	DivingFishDeveloperToken string `json:"diving_fish_developer_token"`

	// data source
	OngeDatasource *DatasourceConf `json:"onge_datasource"`
}

//go:embed templates/config.json
var defaultConfBytes []byte

func ReadConfig(configPath string) *Config {

	configData, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Error("read config file failed. generating default file...")
			confFile, _ := os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
			confFile.Write(defaultConfBytes)
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
