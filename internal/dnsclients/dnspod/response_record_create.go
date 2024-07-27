// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package dnspod

type RecordCreateResponse struct {
	BaseResponse

	Record struct {
		Id     any    `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"record"`
}
