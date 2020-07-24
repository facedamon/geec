package lru

import (
	"reflect"
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestCache_Get(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestCache_Add(t *testing.T) {
	tests := []struct {
		key   string
		value Value
	}{
		{"k1", String("v1")},
		{"k2", String("v2")},
		{"k3", String("v3")},
	}
	ks := make([]string, 0)
	lru := New(int64(len("k1v1k2v2")), func(key string, value Value) {
		ks = append(ks, key)
	})
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			lru.Add(tt.key, tt.value)
		})
	}

	if _, ok := lru.Get("k1"); ok || lru.Len() != 2 {
		t.Fatalf("Add key1 failed")
	}

	if !reflect.DeepEqual(ks, []string{"k1"}) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", "k1")
	}
}
