package service

import (
	"testing"
)

func Test_Cache(t *testing.T) {

	t.Run("test 1", func(t *testing.T) {
		service := NewLRUCacheService(10)
		service.Put("1", "obj1")
		value, ok := service.Get("1")
		want := "obj1"
		if !ok {
			t.Error("test1 failed ")
		}
		if value.(string) != want {
			t.Error("test1 failed ")
		}
	})
}

func Test2_Cache(t *testing.T) {
	t.Run("test 2", func(t *testing.T) {
		service := NewLRUCacheService(10)
		// service.Put("1", "obj1")
		value, ok := service.Get("1")
		want := "obj1"
		if !ok {
			t.Error("test2 failed ")
		} else if value.(string) != want {
			t.Error("test2 failed ")
		}
	})
}

func Test3_Cache(t *testing.T) {
	t.Run("test eviction", func(t *testing.T) {
		service := NewLRUCacheService(4)
		service.Put("1", "obj1")
		service.Put("2", "obj2")
		service.Put("3", "obj3")
		service.Put("4", "obj4")
		service.Put("5", "obj5")
		_, ok := service.Get("1")
		if !ok {
		} else {
			t.Error("test eviction got obj 1 failed ")
		}

		got2, ok := service.Get("2")
		want2 := "obj2"
		if !ok {
			t.Error("test eviction get obj2 failed ")
		} else if got2.(string) != want2 {
			t.Error("test eviction get obj2 failed ")
		}

		got5, ok := service.Get("5")
		want5 := "obj5"
		if !ok {
			t.Error("test eviction get obj5 failed ")
		} else if got5.(string) != want5 {
			t.Error("test eviction get obj5 failed ")
		}
	})
}
