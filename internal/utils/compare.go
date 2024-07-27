// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package utils

import (
	"bytes"
	"encoding/json"
)

// EqualConfig 使用JSON对比配置
func EqualConfig(config1 any, config2 any) bool {
	config1JSON, _ := json.Marshal(config1)
	config2JSON, _ := json.Marshal(config2)
	return bytes.Equal(config1JSON, config2JSON)
}
