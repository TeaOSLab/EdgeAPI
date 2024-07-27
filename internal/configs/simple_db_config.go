// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package configs

import (
	"fmt"
	"net/url"
	"os"

	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/dbs"
	"gopkg.in/yaml.v3"
)

type SimpleDBConfig struct {
	User       string   `yaml:"user"`
	Password   string   `yaml:"password"`
	Database   string   `yaml:"database"`
	Host       string   `yaml:"host"`
	BoolFields []string `yaml:"boolFields,omitempty"`
}

func ParseSimpleDBConfig(data []byte) (*SimpleDBConfig, error) {
	var config = &SimpleDBConfig{}
	err := yaml.Unmarshal(data, config)
	return config, err
}

func (this *SimpleDBConfig) GenerateOldConfig() error {
	var dbConfig = &dbs.DBConfig{
		Driver: "mysql",
		Dsn:    url.QueryEscape(this.User) + ":" + this.Password + "@tcp(" + this.Host + ")/" + url.PathEscape(this.Database) + "?charset=utf8mb4&timeout=30s&multiStatements=true",
		Prefix: "edge",
	}
	dbConfig.Models.Package = "internal/db/models"

	var config = &dbs.Config{
		DBs: map[string]*dbs.DBConfig{
			Tea.Env: dbConfig,
		},
	}
	config.Default.DB = Tea.Env
	config.Fields = map[string][]string{
		"bool": this.BoolFields,
	}

	oldConfigYAML, encodeErr := yaml.Marshal(config)
	if encodeErr != nil {
		return encodeErr
	}

	var targetFile = Tea.ConfigFile(".db.yaml")
	err := os.WriteFile(targetFile, oldConfigYAML, 0666)
	if err != nil {
		return fmt.Errorf("create database config file failed: %w", err)
	}

	return nil
}
