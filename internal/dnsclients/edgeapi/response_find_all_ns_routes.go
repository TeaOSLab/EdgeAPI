// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package edgeapi

type FindAllNSRoutesResponse struct {
	BaseResponse

	Data struct {
		NSRoutes []struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"nsRoutes"`
	} `json:"data"`
}
