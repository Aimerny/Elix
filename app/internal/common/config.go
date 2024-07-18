package common

import (
	_ "embed"
	"errors"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"os"
)

var elixConfig *Config

type Config struct {
	BotToken string `json:"token"`
	Compress bool   `json:"compress"`

	//==== global part =====
	ApiServerPort     int    `json:"api_server_port"`
	WsProxyServerPort int    `json:"ws_proxy_server_port"`
	LogLevel          string `json:"log_level"`
	DataDirPath       string `json:"data_dir_path"`

	//==== onge module ====
	OngeEnable               bool            `json:"onge_enable"`
	DivingFishDeveloperToken string          `json:"diving_fish_developer_token"`
	OngeDatasource           *DatasourceConf `json:"onge_datasource"`
}

//go:embed templates/config.json
var defaultConfBytes []byte

func GlobalConf() *Config {
	if elixConfig == nil {
		elixConfig = ReadConfig(*ConfigPathParam)
	}
	return elixConfig
}

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

func SaveConfig(conf *Config) error {
	if conf == nil {
		return errors.New("config is nil")
	}
	confBytes, err := jsoniter.Marshal(conf)
	if err != nil {
		return err
	}
	confFile, err := os.OpenFile(*ConfigPathParam, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	defer confFile.Close()
	if err != nil {
		return errors.New("open config file failed, err: " + err.Error())
	}
	_, err = confFile.Write(confBytes)
	return nil
}
