// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package huaweidns

type CustomLinesResponse struct {
	Lines []struct {
		LineId string `json:"line_id"`
		Name   string `json:"name"`
	} `json:"lines"`
}
