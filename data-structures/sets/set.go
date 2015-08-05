package set

type Set struct {
	cardinality int
	elements    []interface{}
}

func NewSet(capacity int) *Set {
	set := &Set{}
	if capacity < 0 {
		capacity = 0
	}
	set.elements = make([]interface{}, 0, capacity)
	return set
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
	set.elements = append(set.elements, value)
	set.cardinality++
}

func (set *Set) Contains(value interface{}) bool {
	return set.indexOf(value) != -1
}

func (set *Set) indexOf(value interface{}) int {
	for i, v := range set.elements {
		if v == value {
			return i
		}
	}
	return -1
}

func (set *Set) Delete(value interface{}) {
	last_index := len(set.elements) - 1
	last_value := set.elements[last_index]
	set.elements[set.indexOf(value)] = last_value
	set.elements = set.elements[:last_index]
	set.cardinality--
}

func (set *Set) Union(other *Set) *Set {
	result := NewSet(0)
	for _, v := range set.elements {
		result.Add(v)
	}
	for _, v := range other.elements {
		result.Add(v)
	}
	return result
}

func (set *Set) Intersection(other *Set) *Set {
	result := NewSet(0)
	var smaller, larger *Set
	if set.cardinality < other.Cardinality() {
		smaller = set
		larger = other
	} else {
		smaller = other
		larger = set
	}
	for _, v := range smaller.elements {
		if larger.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (set *Set) Difference(other *Set) *Set {
	result := NewSet(0)
	for _, v := range set.elements {
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
