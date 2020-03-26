package main

type stringList struct {
	list  []string
	count int
}

func (s *stringList) hasNext() bool {
	return len(s.list)-1 >= s.count
}

func (s *stringList) next() {
	s.count++
}

func (s *stringList) getItem() string {
	return s.list[s.count]
}
