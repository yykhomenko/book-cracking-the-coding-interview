package p9ch1

import (
	"fmt"
	"testing"
)

func TestHashtable_Add(t *testing.T) {
	h := New()
	h.Add("one", 1)
	h.Add("five", 5)
	fmt.Println(h)
}
