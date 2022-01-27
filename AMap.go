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

// NewAMap returns a new AMap
func NewAMap() *AMap {
	return &AMap{
		Map: make(map[string]interface{}),
		M:   &sync.Mutex{},
	}

}

// Get a value from the map
func (a *AMap) Get(key string) interface{} {
	a.M.Lock()
	defer a.M.Unlock()
	return a.Map[key]
}

// Set a key to a value
func (a *AMap) Set(key string, value interface{}) {
	a.M.Lock()
	defer a.M.Unlock()
	a.Map[key] = value
}

// Delete a key from the map
func (a *AMap) Delete(key string) {
	a.M.Lock()
	defer a.M.Unlock()
	delete(a.Map, key)
}

// Returns true if the key in the map
func (a *AMap) ContainsKey(key string) bool {
	a.M.Lock()
	defer a.M.Unlock()
	_, ok := a.Map[key]
	return ok
}

// Returns true if the value in the map
func (a *AMap) ContainsValue(value interface{}) bool {
	a.M.Lock()
	defer a.M.Unlock()
	for _, v := range a.Map {
		if v == value {
			return true
		}
	}
	return false
}

// Returns a slice of the keys in the map
func (a *AMap) Keys() []string {
	a.M.Lock()
	defer a.M.Unlock()
	keys := make([]string, 0)
	for k := range a.Map {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of the values in the map
func (a *AMap) Values() []interface{} {
	a.M.Lock()
	defer a.M.Unlock()
	values := make([]interface{}, 0)
	for _, v := range a.Map {
		values = append(values, v)
	}
	return values
}

// Returns the size of the map
func (a *AMap) Size() int {
	return len(a.Map)
}

// Clear the map
func (a *AMap) Clear() {
	a.M.Lock()
	defer a.M.Unlock()
	a.Map = make(map[string]interface{})
}

// Prints the map to the console
func (a *AMap) ToString() string {
	return fmt.Sprintf("%v", a.Map)
}

// ToJSON returns the map as a JSON string
func (a *AMap) ToJSON() string {
	return fmt.Sprintf("%v", a.Map)
}

// Get a random key from the map
func GetRandomKey(a *AMap) string {
	a.M.Lock()
	defer a.M.Unlock()
	keys := a.Keys()
	return keys[rand.Intn(len(keys))]
}

// Get a random value from the map
func GetRandomValue(a *AMap) interface{} {
	a.M.Lock()
	defer a.M.Unlock()
	values := a.Values()
	return values[rand.Intn(len(values))]
}

// Prints the map to the console
func (a *AMap) Print() {
	fmt.Println(a.Map)
}
