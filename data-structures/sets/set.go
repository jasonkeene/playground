package set

type Set struct {
	elements map[interface{}]bool
}

func NewSet() *Set {
	return &Set{elements: make(map[interface{}]bool)}
}

func (set *Set) visit(visitor func(interface{})) {
	for v, _ := range set.elements {
		visitor(v)
	}
}

func (set *Set) IsEmpty() bool {
	return set.Cardinality() <= 0
}

func (set *Set) Cardinality() int {
	return len(set.elements)
}

func (set *Set) Add(value interface{}) {
	set.elements[value] = true
}

func (set *Set) take(predicate func(interface{}) bool) *Set {
	result := NewSet()
	set.visit(func(v interface{}) {
		if predicate(v) {
			result.Add(v)
		}
	})
	return result
}

func (set *Set) Contains(value interface{}) bool {
	_, found := set.elements[value]
	return found
}

func (set *Set) Delete(value interface{}) {
	delete(set.elements, value)
}

func (set *Set) Clear() {
	set.elements = make(map[interface{}]bool)
}

func (set *Set) Union(other *Set) *Set {
	result := NewSet()
	set.visit(func(v interface{}) { result.Add(v) })
	other.visit(func(v interface{}) { result.Add(v) })
	return result
}

func (set *Set) Intersection(other *Set) *Set {
	if set.Cardinality() < other.Cardinality() {
		return set.take(func(v interface{}) bool { return other.Contains(v) })
	} else {
		return other.take(func(v interface{}) bool { return set.Contains(v) })
	}
}

func (set *Set) Difference(other *Set) *Set {
	return set.take(func(v interface{}) bool { return !other.Contains(v) })
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
