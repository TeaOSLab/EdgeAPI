// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package configs

import (
	"errors"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/dbs"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadDBConfig() (*dbs.Config, error) {
	var config = &dbs.Config{}
	for _, filename := range []string{".db.yaml", "db.yaml"} {
		configData, err := os.ReadFile(Tea.ConfigFile(filename))
		if err != nil {
			continue
		}
		err = yaml.Unmarshal(configData, config)
		return config, err
	}

	return nil, errors.New("could not find database config file '.db.yaml' or 'db.yaml'")
}
