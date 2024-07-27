// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package regexputils_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeAPI/internal/utils/regexputils"
	"github.com/iwind/TeaGo/assert"
)

func TestExpr(t *testing.T) {
	var a = assert.NewAssertion(t)

	a.IsTrue(regexputils.YYYYMMDD.MatchString("20221011"))
	a.IsFalse(regexputils.YYYYMMDD.MatchString("202210"))

	a.IsTrue(regexputils.YYYYMM.MatchString("202210"))
	a.IsFalse(regexputils.YYYYMM.MatchString("20221011"))
}
