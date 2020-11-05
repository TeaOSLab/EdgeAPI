package models

//
type RegionCity struct {
	Id         uint32 `field:"id"`         // ID
	ProvinceId uint32 `field:"provinceId"` // 省份ID
	Name       string `field:"name"`       // 名称
	Codes      string `field:"codes"`      // 代号
	State      uint8  `field:"state"`      // 状态
}

type RegionCityOperator struct {
	Id         interface{} // ID
	ProvinceId interface{} // 省份ID
	Name       interface{} // 名称
	Codes      interface{} // 代号
	State      interface{} // 状态
}

func NewRegionCityOperator() *RegionCityOperator {
	return &RegionCityOperator{}
}
