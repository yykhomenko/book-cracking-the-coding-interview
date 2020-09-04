package hashtable

import (
	"strconv"
)

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
	h.array[index] = v
}

func hash(v string) int {
	return len(v)
}

func (h Hashtable) String() string {
	s := ""
	for k, v := range h.array {
		s += strconv.Itoa(k) + " " + strconv.Itoa(v) + "\n"
	}
	return s
}
