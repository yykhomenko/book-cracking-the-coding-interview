package p9ch1

import "testing"

func TestHashtable_Add(t *testing.T) {
	h := New()
	h.Add("one", 1)
}
