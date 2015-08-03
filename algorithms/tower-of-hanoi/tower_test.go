package tower

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTower(t *testing.T) {
	Convey("given a new tower", t, func() {
		tower := NewTower(8)
		Convey("it starts with all rings on peg A", func() {
			So(tower.Pegs["A"].Values, ShouldResemble, []int{8, 7, 6, 5, 4, 3, 2, 1})
			So(tower.Pegs["B"].Values, ShouldBeEmpty)
			So(tower.Pegs["C"].Values, ShouldBeEmpty)
		})
		Convey("it can move entire tower from A to B", func() {
			tower.Move(8, "A", "B")
			So(tower.Pegs["A"].Values, ShouldBeEmpty)
			So(tower.Pegs["B"].Values, ShouldResemble, []int{8, 7, 6, 5, 4, 3, 2, 1})
			So(tower.Pegs["C"].Values, ShouldBeEmpty)
		})
		Convey("it can move entire tower from A to C", func() {
			tower.Move(8, "A", "C")
			So(tower.Pegs["A"].Values, ShouldBeEmpty)
			So(tower.Pegs["B"].Values, ShouldBeEmpty)
			So(tower.Pegs["C"].Values, ShouldResemble, []int{8, 7, 6, 5, 4, 3, 2, 1})
		})
		Convey("it can move bits of the tower from A to C", func() {
			tower.Move(4, "A", "C")
			So(tower.Pegs["A"].Values, ShouldResemble, []int{8, 7, 6, 5})
			So(tower.Pegs["B"].Values, ShouldBeEmpty)
			So(tower.Pegs["C"].Values, ShouldResemble, []int{4, 3, 2, 1})
		})
	})
	Convey("given a tower that was moved to B", t, func() {
		tower := NewTower(8)
		tower.Move(8, "A", "B")
		Convey("it can move back to A", func() {
			tower.Move(8, "B", "A")
			So(tower.Pegs["A"].Values, ShouldResemble, []int{8, 7, 6, 5, 4, 3, 2, 1})
			So(tower.Pegs["B"].Values, ShouldBeEmpty)
			So(tower.Pegs["C"].Values, ShouldBeEmpty)
		})
		Convey("it can move to C", func() {
			tower.Move(8, "B", "C")
			So(tower.Pegs["A"].Values, ShouldBeEmpty)
			So(tower.Pegs["B"].Values, ShouldBeEmpty)
			So(tower.Pegs["C"].Values, ShouldResemble, []int{8, 7, 6, 5, 4, 3, 2, 1})
		})
	})
	Convey("given a checkable tower", t, func() {
		tower := NewTower(8)
		tower.Check = true
		Convey("it doesn't put larger disk on smaller ones", func() {
			So(func() {
				tower.Move(8, "A", "C")
			}, ShouldNotPanic)
			So(func() {
				tower.Move(8, "C", "B")
			}, ShouldNotPanic)
			So(func() {
				tower.Move(8, "B", "A")
			}, ShouldNotPanic)
		})
		Convey("it panics when forced to put larger disks on smaller ones", func() {
			So(func() {
				tower.Move(4, "A", "C")
				tower.Move(4, "A", "C")
			}, ShouldPanic)
		})
	})
}
