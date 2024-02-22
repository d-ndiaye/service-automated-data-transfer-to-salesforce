package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	FolderPath   string `yaml:"folderPath"`
	BackupFolder string `yaml:"backupFolder"`
	Mysql        Mysql  `yaml:"mysql"`
}

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

func Load(file string) (error, Config) {
	var config Config
	_, err := os.Stat(file)
	if err != nil {
		return err, config
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("can't load config file: %s", err.Error()), config
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("can't unmarshal config file: %s", err.Error()), config
	}
	return nil, config
}
