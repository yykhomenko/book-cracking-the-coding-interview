package arraylist

import (
	"fmt"
	"testing"
)

func TestArrayList_String(t *testing.T) {
	l := New()
	fmt.Println(l.String())
}
