// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package dnsclients_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeAPI/internal/dnsclients"
	"github.com/TeaOSLab/EdgeAPI/internal/dnsclients/dnstypes"
	"github.com/TeaOSLab/EdgeAPI/internal/errors"
)

func TestBaseProvider_WrapError(t *testing.T) {
	var provider = &dnsclients.BaseProvider{}
	t.Log(provider.WrapError(nil, "example.com", &dnstypes.Record{
		Id:    "",
		Name:  "a",
		Type:  "A",
		Value: "192.168.1.100",
		Route: "",
		TTL:   3600,
	}))
	t.Log(provider.WrapError(errors.New("fake error"), "example.com", &dnstypes.Record{
		Id:    "",
		Name:  "a",
		Type:  "A",
		Value: "192.168.1.100",
		Route: "",
		TTL:   3600,
	}))
	t.Log(provider.WrapError(errors.New("fake error"), "example.com", &dnstypes.Record{
		Id:    "",
		Name:  "",
		Type:  "A",
		Value: "192.168.1.100",
		Route: "",
		TTL:   3600,
	}))
}
