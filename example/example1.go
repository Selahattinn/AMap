package main

import (
	"fmt"

	amap "github.com/Selahattinn/AMap"
)

func main() {
	amap := amap.NewAMap()
	amap.Set("key1", "value1")
	fmt.Println(amap.Get("key1"))
}
