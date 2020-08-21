package p9ch1

import "fmt"

type Hashtable struct {
	array []int
}

// New returns initialized Hashtable
func New() *Hashtable {
	h := new(Hashtable)
	h.array = make([]int, 10, 10)
	return h
}

// Add the pair
func (h *Hashtable) Add(k string, v int) {
	index := hash(k) % len(h.array)
	fmt.Println(index)
}

func hash(v string) int {
	return len(v)
}
