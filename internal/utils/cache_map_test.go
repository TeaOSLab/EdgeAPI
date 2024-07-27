// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package utils

import (
	"testing"

	"github.com/iwind/TeaGo/assert"
)

func TestNewCacheMap(t *testing.T) {
	var a = assert.NewAssertion(t)

	m := NewCacheMap()
	{
		m.Put("Hello", "World")
		v, ok := m.Get("Hello")
		a.IsTrue(ok)
		a.IsTrue(v == "World")
	}

	{
		v, ok := m.Get("Hello1")
		a.IsFalse(ok)
		a.IsTrue(v == nil)
	}
}
