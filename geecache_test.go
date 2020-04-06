package geecache

import (
	"testing"
	"reflect"
	"log"
	"fmt"
)



func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil 
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Fatal("callback failed")
	}
}


var db = map[string]string {
	"tom" : "630",
	"jack" : "589",
	"sam" : "567",
}


func TestGroupGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))
	gee := NewGroup("scores", 2<<10, GetterFunc(
    	      func(key string) ([]byte, error) {
		      log.Println("SLOWDb search key", key)
		      if v, ok := db[key]; ok {
			      if _, ok := loadCounts[key]; !ok {
				      loadCounts[key] = 0
			      }
			      loadCounts[key] += 1
			      return []byte(v), nil
		      }
		      return nil, fmt.Errorf("%s not exist", key)
	      }))
        for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatal("get failed")
		}
		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}
	if _, err := gee.Get("unknown"); err == nil {
		t.Fatal("hit unexpected")
	}
}
