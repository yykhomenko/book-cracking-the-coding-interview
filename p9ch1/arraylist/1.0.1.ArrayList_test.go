package arraylist

import (
	"fmt"
	"testing"
)

func TestArrayList_String(t *testing.T) {
	l := New()
	l.Add(3)
	l.Add(5)
	fmt.Println(l.String())
}
