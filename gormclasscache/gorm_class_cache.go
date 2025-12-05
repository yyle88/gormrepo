package gormclasscache

import (
	"sync"

	"github.com/yyle88/erero"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/mutexmap"
	"github.com/yyle88/mutexmap/cachemap"
)

// UmcV1 提供缓存功能，避免每次都调用 Columns 函数生成 gormcnm 对象，提升性能
// UmcV1 provides caching feature, avoiding the need to invoke the Columns function to generate gormcnm objects each time, thus improving performance.
func UmcV1[MOD gormclass.ModelClass[CLS], CLS any](a MOD, cache *cachemap.Map[string, interface{}]) (MOD, CLS) {
	vax, _ := cache.Getset(a.TableName(), func() (interface{}, error) {
		return a.Columns(), nil
	})
	cls, ok := vax.(CLS)
	if !ok {
		panic(erero.Errorf("wrong TABLE_NAME=%s", a.TableName()))
	}
	return a, cls
}

// UmcV2 提供缓存功能，避免每次都调用 Columns 函数生成 gormcnm 对象，提升性能
// UmcV2 provides caching feature, avoiding the need to invoke the Columns function to generate gormcnm objects each time, thus improving performance.
func UmcV2[MOD gormclass.ModelClass[CLS], CLS any](a MOD, cache *mutexmap.Map[string, interface{}]) (MOD, CLS) {
	vax, _ := cache.Getset(a.TableName(), func() interface{} {
		return a.Columns()
	})
	cls, ok := vax.(CLS)
	if !ok {
		panic(erero.Errorf("wrong TABLE_NAME=%s", a.TableName()))
	}
	return a, cls
}

// UmcV3 提供缓存功能，避免每次都调用 Columns 函数生成 gormcnm 对象，提升性能
// UmcV3 provides caching feature, avoiding the need to invoke the Columns function to generate gormcnm objects each time, thus improving performance.
func UmcV3[MOD gormclass.ModelClass[CLS], CLS any](a MOD, cache *sync.Map) (MOD, CLS) {
	value, ok := cache.Load(a.TableName())
	if !ok {
		value = a.Columns()
		cache.Store(a.TableName(), value) //这里在并发时有可能多次存储，但是这并不影响任何功能，因此忽略这个情况
	}
	cls, ok := value.(CLS)
	if !ok {
		panic(erero.Errorf("wrong TABLE_NAME=%s", a.TableName()))
	}
	return a, cls
}
