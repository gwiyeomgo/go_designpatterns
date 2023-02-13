package main

type set map[interface{}]struct{}

func (s set) Add(v interface{}) {
	s[v] = struct{}{}
}

func (s set) Remove(v interface{}) {
	delete(s, v)
}

func (s set) Contain(v interface{}) bool {
	_, ok := s[v]
	return ok
}

func (s set) Length() int {
	return len(s)
}
