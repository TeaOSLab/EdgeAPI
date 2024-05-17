// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package dnspod

type ResponseInterface interface {
	IsOk() bool
	LastError() (code string, message string)
}
