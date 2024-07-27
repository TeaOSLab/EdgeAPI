// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package dnspod

type RecordLineResponse struct {
	BaseResponse

	Lines []string `json:"lines"`
}
