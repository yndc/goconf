package utils

// Set is an implementation of a hash set with strings as keys
type Set struct {
	set map[string]struct{}
}

// Add adds keys to the set
func (s *Set) Add(keys ...string) {
	for _, k := range keys {
		if _, ok := s.set[k]; !ok {
			s.set[k] = struct{}{}
		}
	}
}

// Remove removes keys from the set, if a key is not found then it's a no-op
func (s *Set) Remove(keys ...string) {
	for _, k := range keys {
		delete(s.set, k)
	}
}

// Get the count of the set
func (s *Set) Count() int {
	return len(s.set)
}

// Merge from another set
func (s *Set) Merge(set *Set) {
	s.Add(set.Values()...)
}

// Has checks if the set has any of the given keys
func (s *Set) Has(key ...string) bool {
	for _, k := range key {
		if _, ok := s.set[k]; ok {
			return true
		}
	}
	return false
}

// HasAll checks if the set has all of the given keys
func (s *Set) HasAll(key ...string) bool {
	for _, k := range key {
		if _, ok := s.set[k]; ok == false {
			return false
		}
	}
	return true
}

// HasOnly checks if the set has only of the given keys
func (s *Set) HasOnly(key ...string) bool {
	otherSet := NewSetWithValues(key...)
	for k := range s.set {
		if !otherSet.Has(k) {
			return false
		}
	}
	return true
}

// NewSetWithValues creates a new set with the given values
func NewSetWithValues(keys ...string) *Set {
	set := &Set{}
	set.set = make(map[string]struct{}, len(keys))
	set.Add(keys...)
	return set
}

// NewSet creates a new set
func NewSet(length int) *Set {
	set := &Set{}
	set.set = make(map[string]struct{}, length)
	return set
}

// Values get all values stored in the set as an array of strings
func (s *Set) Values() []string {
	values := make([]string, len(s.set))
	i := 0
	for k := range s.set {
		values[i] = k
		i++
	}
	return values
}
