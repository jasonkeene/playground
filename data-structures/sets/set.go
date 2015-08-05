package set

type Set struct {
	cardinality int
	elements    map[interface{}]bool
}

func NewSet() *Set {
	return &Set{elements: make(map[interface{}]bool)}
}

func (set *Set) IsEmpty() bool {
	return set.cardinality <= 0
}

func (set *Set) Cardinality() int {
	return set.cardinality
}

func (set *Set) Add(value interface{}) {
	if set.Contains(value) {
		return
	}
	set.elements[value] = true
	set.cardinality++
}

func (set *Set) Contains(value interface{}) bool {
	_, found := set.elements[value]
	return found
}

func (set *Set) Delete(value interface{}) {
	delete(set.elements, value)
	set.cardinality--
}

func (set *Set) Clear() {
	set.elements = make(map[interface{}]bool)
	set.cardinality = 0
}

func (set *Set) Union(other *Set) *Set {
	result := NewSet()
	for v, _ := range set.elements {
		result.Add(v)
	}
	for v, _ := range other.elements {
		result.Add(v)
	}
	return result
}

func (set *Set) Intersection(other *Set) *Set {
	result := NewSet()
	var smaller, larger *Set
	if set.cardinality < other.Cardinality() {
		smaller = set
		larger = other
	} else {
		smaller = other
		larger = set
	}
	for v, _ := range smaller.elements {
		if larger.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (set *Set) Difference(other *Set) *Set {
	result := NewSet()
	for v, _ := range set.elements {
		if !other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (set *Set) SymmetricDifference(other *Set) *Set {
	return set.Difference(other).Union(other.Difference(set))
}

func (set *Set) IsSubset(other *Set) bool {
	return set.Difference(other).IsEmpty()
}

func (set *Set) IsSuperset(other *Set) bool {
	return other.IsSubset(set)
}
