package hanoi

import (
	"errors"
	"fmt"
	"log"
)

type Peg struct {
	Values []int
}

func NewPeg(d []int) *Peg {
	return &Peg{
		Values: d,
	}
}

func (peg *Peg) Pop() int {
	length := len(peg.Values)
	if length == 0 {
		panic("popping from empty peg")
	}
	value := peg.Values[length-1]
	peg.Values = peg.Values[0 : length-1]
	return value
}

func (peg *Peg) Push(value int) {
	peg.Values = append(peg.Values, value)
}

func (peg Peg) Check() {
	if len(peg.Values) == 0 {
		return
	}
	value := peg.Values[0]
	for i := range peg.Values[1:] {
		if peg.Values[i] > value {
			panic("peg contains disks that are out of order")
		}
		value = peg.Values[i]
	}
}

type Tower struct {
	Pegs  map[string]*Peg
	Check bool
}

func NewTower(size int) *Tower {
	tower := &Tower{
		Pegs: map[string]*Peg{
			"A": NewPeg(make([]int, 0, size)),
			"B": NewPeg(make([]int, 0, size)),
			"C": NewPeg(make([]int, 0, size)),
		},
	}
	for i := size; i > 0; i-- {
		tower.Pegs["A"].Push(i)
	}
	return tower
}

func (tower *Tower) Move(count int, from, to string) {
	from_peg := tower.Pegs[from]
	to_peg := tower.Pegs[to]
	other, err := lookupOther(from, to)
	if err != nil {
		log.Fatal(err)
	}
	other_peg := tower.Pegs[other]

	hanoi(count, from_peg, to_peg, other_peg, tower.Check)
}

func hanoi(count int, from, to, other *Peg, check bool) {
	if count > 1 {
		// move sub problem from "from" to "other"
		hanoi(count-1, from, other, to, check)
	}

	// work is done here
	value := from.Pop()
	to.Push(value)

	// optionally check state of pegs
	if check {
		from.Check()
		to.Check()
		other.Check()
	}

	if count > 1 {
		// move sub problem from "other" to "to"
		hanoi(count-1, other, to, from, check)
	}
}

func lookupOther(from, to string) (string, error) {
	peg_names := map[string]bool{
		"A": true,
		"B": true,
		"C": true,
	}
	delete(peg_names, from)
	delete(peg_names, to)
	if len(peg_names) != 1 {
		return "", fmt.Errorf("bad length for peg_names: %d", len(peg_names))
	}
	for k := range peg_names {
		return k, nil
	}
	return "", errors.New("no peg_name available")
}
