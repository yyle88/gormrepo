package example2models

import (
	"sync"

	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/gormclasscache"
	"github.com/yyle88/mutexmap"
	"github.com/yyle88/mutexmap/cachemap"
)

// 这个全局变量的最佳实践，就是和 models 数据放在一起，而里面的数量就是填写数据表的（随意的预估）个数，当然也可以填0
var cache1 = cachemap.NewMap[string, interface{}](2)

// UmcV1 就是用的工具包中的 UmcV1 逻辑，由于go目前不允许类的成员函数名为泛型，这里只能是定义个普通的函数，函数内部用全局变量缓存信息
func UmcV1[MOD gormclass.ClassType[CLS], CLS any](a MOD) (MOD, CLS) {
	return gormclasscache.UmcV1(a, cache1)
}

// 这个全局变量的最佳实践，就是和 models 数据放在一起，而里面的数量就是填写数据表的（随意的预估）个数，当然也可以填0
var cache2 = mutexmap.NewMap[string, interface{}](2)

// UmcV2 就是用的工具包中的 UmcV2 逻辑，由于go目前不允许类的成员函数名为泛型，这里只能是定义个普通的函数，函数内部用全局变量缓存信息
func UmcV2[MOD gormclass.ClassType[CLS], CLS any](a MOD) (MOD, CLS) {
	return gormclasscache.UmcV2(a, cache2)
}

// 这是第二套方案的缓存变量
var cache3 = &sync.Map{}

// UmcV3 就是用的工具包中的 UmcV3 逻辑
func UmcV3[MOD gormclass.ClassType[CLS], CLS any](a MOD) (MOD, CLS) {
	return gormclasscache.UmcV3(a, cache3)
}
