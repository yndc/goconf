package utils

import "strings"

type Path struct {
	value []string
}

func NewPath() *Path {
	return &Path{
		value: make([]string, 0),
	}
}

func (p *Path) Copy() *Path {
	new := NewPath()
	copy(p.value, new.value)
	return new
}

func Parse(source string) []string {
	splitted := strings.Split(source, ".")
	if len(splitted) > 0 {
		return splitted
	}
	return nil
}

func (p *Path) Add(path ...string) {
	p.value = append(p.value, path...)
}

func (p *Path) Back(count int) {
	length := len(p.value)
	if length >= count && count > 0 {
		p.value = p.value[:len(p.value)-count]
	}
}
