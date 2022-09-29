package main

import (
	"fmt"
	"time"

	"github.com/muesli/cache2go"
)

type myStruct struct {
	title   string
	content interface{}
}

func main() {
	table := "test"
	key := "oscome"

	// 新建 CacheTable
	cache := cache2go.Cache(table)
	setDataLoader(table)
	Callback(table, key)

	val := myStruct{"what", "hello oscome"}

	// 添加一个kv，过期时间 2s
	cache.Add(key, 2*time.Second, &val)

	// 获取value
	res, err := cache.Value(key)
	if err == nil {
		fmt.Printf("Found value=%+v", res.Data().(*myStruct))
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	// sleep 3s
	time.Sleep(3 * time.Second)
	res, err = cache.Value(key)
	if err != nil {
		fmt.Println("Item is not cached (anymore).")
	} else {
		fmt.Println("cache", res)
	}

	cache.Add(key, 0, &val)

	// 增加删除回调函数
	cache.SetAboutToDeleteItemCallback(func(e *cache2go.CacheItem) {
		fmt.Println("Deleted:", e.Key(), e.Data().(*myStruct).title, e.CreatedOn())
	})

	fmt.Println("to delete", key)
	// 删除
	cache.Delete(key)

	// 清除所有item
	cache.Flush()

}

func setDataLoader(table string) {
	cache := cache2go.Cache(table)
	// 在尝试访问不存在的key时将调用该回调
	cache.SetDataLoader(func(key interface{}, args ...interface{}) *cache2go.CacheItem {
		val := "this is " + key.(string) + ", xiaoou"
		fmt.Println("SetDataLoader", key, "--", val)
		return cache2go.NewCacheItem(key, 0, val)
	})
}

func Callback(table, key string) {
	cache := cache2go.Cache(table)

	// 设置添加回调
	cache.SetAddedItemCallback(func(entry *cache2go.CacheItem) {
		fmt.Println("callback 1:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	// 可以设置多个，但会被覆盖
	cache.SetAddedItemCallback(func(entry *cache2go.CacheItem) {
		fmt.Println("callback 2:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	cache.SetAboutToDeleteItemCallback(func(entry *cache2go.CacheItem) {
		fmt.Println("del 1:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	// 可以设置多个，但会被覆盖
	cache.SetAboutToDeleteItemCallback(func(entry *cache2go.CacheItem) {
		fmt.Println("del 2:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	cache.Add(key, 0, "xiaoou")

	res, err := cache.Value(key)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	cache.Delete(key)

	fmt.Println(key, res.Data())

	cache.Delete(key)

	cache.RemoveAddedItemCallbacks()

	res = cache.Add(key+"2", time.Second, "oscome")

	res.SetAboutToExpireCallback(func(i interface{}) {
		fmt.Println("expire", i.(string))
	})

	time.Sleep(2 * time.Second)
}
