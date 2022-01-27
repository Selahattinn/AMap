package amap

import (
	"reflect"
	"sync"
	"testing"
)

func generateMap() *AMap {
	amap := NewAMap()
	return amap
}

func TestAMap_Get(t *testing.T) {
	amap := generateMap()
	amap.Map["key"] = "value"
	type fields struct {
		Map map[string]interface{}
		M   *sync.Mutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{"Test for success ", fields{Map: amap.Map, M: amap.M}, args{"key"}, "value"}, {"Test for fail ", fields{Map: amap.Map, M: amap.M}, args{"key1"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AMap{
				Map: tt.fields.Map,
				M:   tt.fields.M,
			}
			if got := a.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
