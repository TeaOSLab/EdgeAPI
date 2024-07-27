// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package dnspod

type DomainInfoResponse struct {
	BaseResponse

	Domain struct {
		Id    any    `json:"id"`
		Name  string `json:"name"`
		Grade string `json:"grade"`
	} `json:"domain"`
}
