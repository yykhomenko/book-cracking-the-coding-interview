package arraylist

import "fmt"

type ArrayList struct {
	raw []int
}

func New() *ArrayList {
	l := new(ArrayList)
	l.raw = make([]int, 5)
	return l
}

func (l ArrayList) String() string {
	return fmt.Sprint(l.raw)
}
