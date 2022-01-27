package amap

import (
	"fmt"
	"math/rand"
	"sync"
)

type AMap struct {
	Map map[string]interface{}
	M   *sync.Mutex
}

func NewAMap() *AMap {
	return &AMap{
		Map: make(map[string]interface{}),
	}

}

func (a *AMap) Get(key string) interface{} {
	a.M.Lock()
	defer a.M.Unlock()
	return a.Map[key]
}
func (a *AMap) Set(key string, value interface{}) {
	a.M.Lock()
	defer a.M.Unlock()
	a.Map[key] = value
}

func (a *AMap) Delete(key string) {
	a.M.Lock()
	defer a.M.Unlock()
	delete(a.Map, key)
}
func (a *AMap) Has(key string) bool {
	a.M.Lock()
	defer a.M.Unlock()
	_, ok := a.Map[key]
	return ok
}
func (a *AMap) Keys() []string {
	a.M.Lock()
	defer a.M.Unlock()
	keys := make([]string, 0)
	for k := range a.Map {
		keys = append(keys, k)
	}
	return keys
}
func (a *AMap) Values() []interface{} {
	a.M.Lock()
	defer a.M.Unlock()
	values := make([]interface{}, 0)
	for _, v := range a.Map {
		values = append(values, v)
	}
	return values
}
func (a *AMap) Size() int {
	return len(a.Map)
}
func (a *AMap) Clear() {
	a.M.Lock()
	defer a.M.Unlock()
	a.Map = make(map[string]interface{})
}
func (a *AMap) ToString() string {
	return fmt.Sprintf("%v", a.Map)
}
func (a *AMap) ToJSON() string {
	return fmt.Sprintf("%v", a.Map)
}
func GetRandomKey(a *AMap) string {
	a.M.Lock()
	defer a.M.Unlock()
	keys := a.Keys()
	return keys[rand.Intn(len(keys))]
}
func GetRandomValue(a *AMap) interface{} {
	a.M.Lock()
	defer a.M.Unlock()
	values := a.Values()
	return values[rand.Intn(len(values))]
}
func (a *AMap) Print() {
	fmt.Println(a.Map)
}
