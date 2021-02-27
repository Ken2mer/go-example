package sync

import (
	"fmt"
)

func ExampleKeyValue() {
	kv := NewKeyValue()
	kv.Set("key", "value")
	value, ok := kv.Get("key")
	if ok {
		fmt.Println(value)
	}
	// Output: value
}
