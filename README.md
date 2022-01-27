# Selahattinn/AMap

Package `Selahattinn/AMap` implements a map with mutex.

The name AMAP stands for "A Map". Like the standard `MAP` but it has some diffrence.

Avaliable Methods:
* Get: Get a value from the map
* Set: Set a key to a value
* Clear: Clear the map
* Delete: Delete a key from the map
* ContainsKey: Returns true if the key in the map
* ContainsValue: Returns true if the value in the map
* Keys: Returns a slice of the keys in the map
* Values: Values returns a slice of the values in the map
* Size: Returns the size of the map
* ToString: Prints the map to the consoles
* ToJSON: ToJSON returns the map as a JSON string
* GetRandomKey: Get a random key from the maps
* GetRandomValue: Get a random value from the map
* Print: Prints the map to the console
---


* [Install](#install)
* [Examples](#examples)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get github.com/Selahattinn/AMap
```

## Examples

Let's start registering a couple of URL paths and handlers:

```go
func main() {
	amap := amap.NewAMap()
	amap.Set("key1", "value1")
	fmt.Println(amap.ContainsKey("key1"))
	fmt.Println(amap.ContainsValue("value1"))
	amap.Delete("key1")
	fmt.Println(amap.ContainsKey("key1"))
	fmt.Println(amap.ContainsValue("value1"))
}
```
```go
func main() {
	amap := amap.NewAMap()
	amap.Set("key1", "value1")
	fmt.Println(amap.ContainsKey("key1"))
    fmt.Println(amap.ContainsValue("value1"))
    amap.Clear()
	fmt.Println(amap.ContainsKey("key1"))
	fmt.Println(amap.ContainsValue("value1"))
}
```

```go
func main() {
	amap := amap.NewAMap()
	amap.Set("key1", "value1")
	fmt.Println(amap.GetRandomKey())
}
```

```go
func main() {
	amap := amap.NewAMap()
	amap.Set("key1", "value1")
	fmt.Println(amap.GetRandomValue())
}
```

```go
func main() {
	amap := amap.NewAMap()
	amap.Set("key1", "value1")
	fmt.Println(amap.Values())
}
```

```go
func main() {
	amap := amap.NewAMap()
	amap.Set("key1", "value1")
	fmt.Println(amap.ToString())
}
```
## License

BSD licensed. See the LICENSE file for details.