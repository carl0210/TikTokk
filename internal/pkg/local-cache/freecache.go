package local_cache

import (
	"github.com/coocood/freecache"
)

var myCache *freecache.Cache

func init() {
	cacheSize := 128 * 1024 * 1024 // 128M
	myCache = freecache.NewCache(cacheSize)
}

// getCacheSize 根据本机内存大小，return合适的cache大小
//func getCacheSize() int {
//	// 获取本机内存大小，以 M 为单位
//	var m runtime.MemStats
//	runtime.ReadMemStats(&m)
//	sz := m.Sys / (1024 * 1024)
//
//}

func SetInLocalCache(key string, value string, expSec int) error {
	return myCache.Set([]byte(key), []byte(value), expSec)
}

func GetInLocalCache(key string) (string, error) {
	value, err := myCache.Get([]byte(key))
	if err != nil {
		return "", err
	}
	return string(value), nil
}

func DelInLocalCache(key string) bool {
	return myCache.Del([]byte(key))
}
