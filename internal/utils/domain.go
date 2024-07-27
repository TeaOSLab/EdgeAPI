// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package utils

import (
	"regexp"
	"strings"
)

var cacheKeyDomainReg1 = regexp.MustCompile(`^(?i)(?:http|https)://([\w-.*]+)`) // 这里支持 *.example.com
var cacheKeyDomainReg2 = regexp.MustCompile(`^([\w-.]+)`)

// ParseDomainFromKey 从Key中获取域名
func ParseDomainFromKey(key string) (domain string) {
	var pieces = cacheKeyDomainReg1.FindStringSubmatch(key)
	if len(pieces) > 1 {
		return strings.ToLower(pieces[1])
	}

	pieces = cacheKeyDomainReg2.FindStringSubmatch(key)
	if len(pieces) > 1 {
		return strings.ToLower(pieces[1])
	}

	return ""
}
