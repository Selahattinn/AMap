package main

import (
	"fmt"

	amap "github.com/Selahattinn/AMap"
)

func main() {
	amap := amap.NewAMap()
	amap.Set("key1", "value1")
	fmt.Println(amap.ContainsKey("key1"))
	fmt.Println(amap.ContainsValue("value1"))
	amap.Delete("key1")
	fmt.Println(amap.ContainsKey("key1"))
	fmt.Println(amap.ContainsValue("value1"))
	fmt.Println(amap.GetRandomKey())
}
