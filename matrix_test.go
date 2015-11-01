package sparse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCSRMatrix(t *testing.T) {
	Convey("Inserted values should be retrieved correctly.", t, func() {
		a := NewCSRMatrix(3, 3)
		a.Set(0, 1, 1.0)
		So(a.Get(0, 1), ShouldEqual, 1.0)

		a.Set(2, 2, 3.0)

		So(a.Get(0, 1), ShouldEqual, 1.0)
		So(a.Get(2, 2), ShouldEqual, 3.0)

		a.Set(1, 0, 2.0)

		So(a.Get(0, 1), ShouldEqual, 1.0)
		So(a.Get(1, 0), ShouldEqual, 2.0)
		So(a.Get(2, 2), ShouldEqual, 3.0)
	})

	Convey("Iterator should yield values correctly.", t, func() {
		a := NewCSRMatrix(3, 3)
		a.Set(0, 1, 1.0)
		a.Set(2, 2, 3.0)
		a.Set(1, 0, 2.0)
		iter := a.IterTriplets()

		triplet, ok := iter.Next()

		So(ok, ShouldEqual, true)
		So(triplet.Row, ShouldEqual, 0)
		So(triplet.Col, ShouldEqual, 1)
		So(triplet.Val, ShouldEqual, 1.0)

		triplet, ok = iter.Next()

		So(ok, ShouldEqual, true)
		So(triplet.Row, ShouldEqual, 1)
		So(triplet.Col, ShouldEqual, 0)
		So(triplet.Val, ShouldEqual, 2.0)

		triplet, ok = iter.Next()

		So(ok, ShouldEqual, true)
		So(triplet.Row, ShouldEqual, 2)
		So(triplet.Col, ShouldEqual, 2)
		So(triplet.Val, ShouldEqual, 3.0)

		triplet, ok = iter.Next()
		So(ok, ShouldEqual, false)
		So(triplet, ShouldEqual, nil)
	})
}

func TestTriplet(t *testing.T) {
	Convey("Triplets should be compared correctly.", t, func() {
		a := &Triplet{0, 1, 1.0}
		b := &Triplet{1, 0, 0.0}
		So(a.LessThan(b), ShouldEqual, true)
		So(b.LessThan(a), ShouldEqual, false)
	})

}

func TestAddCSR(t *testing.T) {
	Convey("Addition of two zero matrices should be a zero matrix.", t, func() {
		a := NewCSRMatrix(2, 2)
		b := NewCSRMatrix(2, 2)
		c := AddCSR(a, b)
		So(c.Get(0, 0), ShouldEqual, 0)
		So(c.Get(0, 1), ShouldEqual, 0)
		So(c.Get(1, 0), ShouldEqual, 0)
		So(c.Get(1, 1), ShouldEqual, 0)
	})
	Convey("Addition of two matrices with no common elements should be correct.", t, func() {
		a := NewCSRMatrix(2, 2)
		a.Set(0, 0, 1.0)
		a.Set(1, 1, 2.0)
		b := NewCSRMatrix(2, 2)
		b.Set(1, 0, 3.0)
		b.Set(0, 1, 4.0)
		c := AddCSR(a, b)
		So(c.Get(0, 0), ShouldEqual, 1.0)
		So(c.Get(0, 1), ShouldEqual, 4.0)
		So(c.Get(1, 0), ShouldEqual, 3.0)
		So(c.Get(1, 1), ShouldEqual, 2.0)
	})
	Convey("Addition of two matrices with common elements should be correct.", t, func() {
		a := NewCSRMatrix(2, 2)
		a.Set(0, 0, 1.0)
		a.Set(1, 1, 2.0)
		a.Set(0, 1, 2.0)
		b := NewCSRMatrix(2, 2)
		b.Set(1, 0, 3.0)
		b.Set(0, 1, 4.0)
		b.Set(1, 1, 1.0)
		c := AddCSR(a, b)
		So(c.Get(0, 0), ShouldEqual, 1.0)
		So(c.Get(0, 1), ShouldEqual, 6.0)
		So(c.Get(1, 0), ShouldEqual, 3.0)
		So(c.Get(1, 1), ShouldEqual, 3.0)
	})
}
