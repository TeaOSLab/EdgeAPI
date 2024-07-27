// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package regexputils

import "regexp"

var (
	YYYYMMDDHH = regexp.MustCompile(`^\d{10}$`)
	YYYYMMDD   = regexp.MustCompile(`^\d{8}$`)
	YYYYMM     = regexp.MustCompile(`^\d{6}$`)
)

var (
	HTTPProtocol = regexp.MustCompile("^(?i)(http|https)://")
)
