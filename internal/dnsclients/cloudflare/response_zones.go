// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package cloudflare

type ZonesResponse struct {
	BaseResponse

	Result []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"result"`
}
