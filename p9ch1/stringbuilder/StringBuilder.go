package stringbuilder

type StringBuilder struct {
	store []byte
}

func (s *StringBuilder) Append(more string) {
	s.store = append(s.store, []byte(more)...)
}

func (s StringBuilder) String() string {
	return string(s.store)
}
