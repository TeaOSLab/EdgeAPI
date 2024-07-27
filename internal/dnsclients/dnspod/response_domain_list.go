// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package dnspod

type DomainListResponse struct {
	BaseResponse

	Info struct {
		DomainTotal int `json:"domain_total"`
		AllTotal    int `json:"all_total"`
		MineTotal   int `json:"mine_total"`
	} `json:"info"`

	Domains []struct {
		Name string `json:"name"`
	} `json:"domains"`
}
