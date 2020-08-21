package p9ch1

import (
	"fmt"
	"testing"

	"github.com/yykhomenko/book-cracking-the-coding-interview/p9ch1/arraylist"
)

func TestHashtable_Add(t *testing.T) {
	h := arraylist.New()
	h.Add("one", 1)
	h.Add("five", 5)
	fmt.Println(h)
}
