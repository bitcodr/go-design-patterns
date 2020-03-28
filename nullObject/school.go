package main

type school struct {
	classrooms []classroom
}

func (s *school) getClass(c classroom) classroom {
	for _, v := range s.classrooms {
		if v.getName() == c.getName() {
			return v
		}
	}
	return &nullClass{}
}

func (s *school) setClass(c ...classroom) {
	s.classrooms = append(s.classrooms, c...)
}
