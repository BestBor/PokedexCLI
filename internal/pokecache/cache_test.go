package pokecache

import (
	"testing"
	"time"
)

func TestCreationOfCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestCacheAdd(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, pcase := range cases {
		cache.Add(pcase.inputKey, pcase.inputVal)
		createdEl, ok := cache.Get(pcase.inputKey)
		if !ok {
			t.Errorf("element: %s not created/found", pcase.inputKey)
			continue
		}
		if string(createdEl) != string(pcase.inputVal) {
			t.Errorf("value: %s mismatch: %s", pcase.inputKey, pcase.inputVal)
			continue
		}
	}

}

func TestReape(t *testing.T) {

	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s - Should not exist", keyOne)
	}

}

func TestReapeFail(t *testing.T) {

	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s - Should not exist", keyOne)
	}

}
