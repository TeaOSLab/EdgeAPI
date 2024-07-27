// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package configs

import (
	"errors"
	"os"

	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/dbs"
	"gopkg.in/yaml.v3"
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
