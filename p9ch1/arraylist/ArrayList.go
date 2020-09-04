package arraylist

import "fmt"

type ArrayList struct {
	raw []int
}

func (l *ArrayList) Add(e int) {
	l.raw = append(l.raw, e)
}

func (l ArrayList) String() string {
	return fmt.Sprint(l.raw)
}
