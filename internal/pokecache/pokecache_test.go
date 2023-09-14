package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Second)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddToCache(t *testing.T) {
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "key1",
			value: []byte("value1"),
		},
		{
			key:   "key2",
			value: []byte("value2"),
		},
		{
			key:   "key3",
			value: []byte("value1"),
		},
		{
			key:   "",
			value: []byte("value3"),
		},
	}
	cache := NewCache(time.Second)
	for _, cs := range cases {
		key := cs.key
		value := cs.value
		cache.Add(key, value)
		actual, ok := cache.Get(key)
		if !ok {
			t.Error("Could not find key-obj added to cache")
		}
		if string(actual) != string(value) {
			t.Errorf("Actual diff from expected: %v vs %v", actual, value)
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + 2*time.Millisecond)
	_, ok := cache.Get(keyOne)
	if ok {
		t.Error("Reap did not occur")
	}
}
func TestFailReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(2 * time.Millisecond)
	_, ok := cache.Get(keyOne)
	if !ok {
		t.Error("Reap occurred when shouldn't")
	}
}
