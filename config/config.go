package config

import (
	"io/ioutil"
	"stchb/logger"

	"github.com/pelletier/go-toml"
)

const File = "configs/config.toml"

type Config struct {
	Bot struct {
		Webhook   string `toml:"webhook"`
		Token     string `toml:"token"`
		Port      string `toml:"port"`
		ChannelID string `toml:"channel_id"`
		Polling   struct {
			Offset  int `toml:"offset"`
			Limit   int `toml:"limit"`
			Timeout int `toml:"timeout"`
		} `toml:"polling"`
		AdministratorPasswordSHA256 string `json:"administrator_password_sha256"`
	} `toml:"bot"`
	DB struct {
		Redis struct {
			Addr     string `toml:"addr"`
			Password string `toml:"password"`
			Db       int    `toml:"db"`
		} `toml:"redis"`
	} `toml:"database"`
	Chat struct {
		Queue int `toml:"queue"`
		Users int `toml:"users"`
	} `toml:"chat"`
}

var config Config

func init() {
	b, err := ioutil.ReadFile(File)
	if err != nil {
		logger.ForWarning(err.Error())
	}

	if err := toml.Unmarshal(b, &config); err != nil {
		logger.ForWarning(err.Error())
	}
}

func GetConfig() Config {
	return config
}
