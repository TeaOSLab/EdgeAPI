// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package edgeapi

type ResponseInterface interface {
	IsValid() bool
	Error() error
}
