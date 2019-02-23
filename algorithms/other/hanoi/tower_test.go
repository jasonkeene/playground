package hanoi_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jasonkeene/playground/algorithms/other/hanoi"
)

func TestTower(t *testing.T) {
	t.Run("new tower", func(t *testing.T) {
		t.Run("all rings start on peg A", func(t *testing.T) {
			tower := hanoi.NewTower(8)

			expected := map[string]*hanoi.Peg{
				"A": hanoi.NewPeg([]int{8, 7, 6, 5, 4, 3, 2, 1}),
				"B": hanoi.NewPeg([]int{}),
				"C": hanoi.NewPeg([]int{}),
			}

			if !cmp.Equal(tower.Pegs, expected) {
				t.Fatal(cmp.Diff(tower.Pegs, expected))
			}
		})

		t.Run("move entire tower from A to B", func(t *testing.T) {
			tower := hanoi.NewTower(8)

			tower.Move(8, "A", "B")

			expected := map[string]*hanoi.Peg{
				"A": hanoi.NewPeg([]int{}),
				"B": hanoi.NewPeg([]int{8, 7, 6, 5, 4, 3, 2, 1}),
				"C": hanoi.NewPeg([]int{}),
			}

			if !cmp.Equal(tower.Pegs, expected) {
				t.Fatal(cmp.Diff(tower.Pegs, expected))
			}
		})

		t.Run("move entire tower from A to C", func(t *testing.T) {
			tower := hanoi.NewTower(8)

			tower.Move(8, "A", "C")

			expected := map[string]*hanoi.Peg{
				"A": hanoi.NewPeg([]int{}),
				"B": hanoi.NewPeg([]int{}),
				"C": hanoi.NewPeg([]int{8, 7, 6, 5, 4, 3, 2, 1}),
			}

			if !cmp.Equal(tower.Pegs, expected) {
				t.Fatal(cmp.Diff(tower.Pegs, expected))
			}
		})

		t.Run("it can move bits of the tower from A to C", func(t *testing.T) {
			tower := hanoi.NewTower(8)

			tower.Move(4, "A", "C")

			expected := map[string]*hanoi.Peg{
				"A": hanoi.NewPeg([]int{8, 7, 6, 5}),
				"B": hanoi.NewPeg([]int{}),
				"C": hanoi.NewPeg([]int{4, 3, 2, 1}),
			}

			if !cmp.Equal(tower.Pegs, expected) {
				t.Fatal(cmp.Diff(tower.Pegs, expected))
			}
		})
	})

	t.Run("everything moved to B", func(t *testing.T) {
		t.Run("move back to A", func(t *testing.T) {
			tower := hanoi.NewTower(8)
			tower.Move(8, "A", "B")

			tower.Move(8, "B", "A")

			expected := map[string]*hanoi.Peg{
				"A": hanoi.NewPeg([]int{8, 7, 6, 5, 4, 3, 2, 1}),
				"B": hanoi.NewPeg([]int{}),
				"C": hanoi.NewPeg([]int{}),
			}

			if !cmp.Equal(tower.Pegs, expected) {
				t.Fatal(cmp.Diff(tower.Pegs, expected))
			}
		})

		t.Run("move to C", func(t *testing.T) {
			tower := hanoi.NewTower(8)
			tower.Move(8, "A", "B")

			tower.Move(8, "B", "C")

			expected := map[string]*hanoi.Peg{
				"A": hanoi.NewPeg([]int{}),
				"B": hanoi.NewPeg([]int{}),
				"C": hanoi.NewPeg([]int{8, 7, 6, 5, 4, 3, 2, 1}),
			}

			if !cmp.Equal(tower.Pegs, expected) {
				t.Fatal(cmp.Diff(tower.Pegs, expected))
			}
		})
	})

	t.Run("checkable tower", func(t *testing.T) {
		t.Run("put larger disk on smaller ones", func(t *testing.T) {
			tower := hanoi.NewTower(8)
			tower.Check = true

			tower.Move(8, "A", "C")
			tower.Move(8, "C", "B")
			tower.Move(8, "B", "A")
		})

		t.Run("panic when forced to put larger disks on smaller ones", func(t *testing.T) {
			tower := hanoi.NewTower(8)
			tower.Check = true
			tower.Move(4, "A", "C")

			defer func() {
				if err := recover(); err == nil {
					t.Fatal("expected move to panic")
				}
			}()

			tower.Move(4, "A", "C")
		})
	})
}
