package dnsclients

import (
	"fmt"
	"github.com/TeaOSLab/EdgeAPI/internal/dnsclients/dnstypes"
)

type BaseProvider struct {
	minTTL int32
}

// WrapError 封装解析相关错误
func (this *BaseProvider) WrapError(err error, domain string, record *dnstypes.Record) error {
	if err == nil {
		return nil
	}

	if record == nil {
		return err
	}

	var fullname string
	if len(record.Name) == 0 {
		fullname = domain
	} else {
		fullname = record.Name + "." + domain
	}
	return fmt.Errorf("record operation failed: '%s %s %s %d': %w", fullname, record.Type, record.Value, record.TTL, err)
}

// SetMinTTL 设置最小TTL
func (this *BaseProvider) SetMinTTL(ttl int32) {
	this.minTTL = ttl
}

// MinTTL 最小TTL
func (this *BaseProvider) MinTTL() int32 {
	if this.minTTL > 0 {
		return this.minTTL
	}
	return 0
}
