// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package utils_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeAPI/internal/utils"
	"github.com/iwind/TeaGo/assert"
)

func TestIsValidMobile(t *testing.T) {
	var a = assert.NewAssertion(t)
	a.IsFalse(utils.IsValidMobile("138"))
	a.IsFalse(utils.IsValidMobile("1382222"))
	a.IsFalse(utils.IsValidMobile("1381234567890"))
	a.IsTrue(utils.IsValidMobile("13812345678"))
}
