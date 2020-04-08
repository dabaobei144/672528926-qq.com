package consistenthash

import (
	"strconv"
	"testing"
	"fmt"
)

func TestHashing(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	hash.Add("6", "4", "2")

	testCases := map[string]string {
		"2" : "2",
		"11" : "2",
		"23" : "4",
		"27" : "2",
	}

	fmt.Println(hash.keys)
	fmt.Println(hash.hashMap)
	for k, v:= range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should has yield %s, but: %s", k, v, hash.Get(k))
		}
	}

	hash.Add("8")
	testCases["27"] = "8" // 验证27所在节点易主, 而其他key无损

	for k, v:= range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should has yield %s", k, v)
		}
	}

}

