package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Database        ConfigDatabase
	Verbose         bool
	Duration        int
	EmojiImagesPath string `yaml:"emoji_images_path"`
}

type ConfigDatabase struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

func ReadConfig(filename string) (cnf Config, err error) {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		return cnf, err
	}

	err = yaml.Unmarshal(buf, &cnf)
	if err != nil {
		return cnf, err
	}
	return cnf, nil
}
