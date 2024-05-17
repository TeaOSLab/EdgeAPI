// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package cloudflare

type ResponseInterface interface {
	IsOk() bool
	LastError() (code int, message string)
}
