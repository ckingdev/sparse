package sparse

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCSRMatrix(t *testing.T) {
	Convey("Inserted values should be retrieved correctly.", t, func() {
		a := NewCSRMatrix(3, 3)
		a.Insert(0, 1, 1.0)
		So(a.getSingleElement(0, 1), ShouldEqual, 1.0)
		fmt.Println(a)
		a.Insert(2, 2, 3.0)
		fmt.Println(a)
		So(a.getSingleElement(0, 1), ShouldEqual, 1.0)
		So(a.getSingleElement(2, 2), ShouldEqual, 3.0)
		a.Insert(1, 0, 2.0)
		fmt.Println(a)
		So(a.getSingleElement(0, 1), ShouldEqual, 1.0)
		So(a.getSingleElement(1, 0), ShouldEqual, 2.0)
		So(a.getSingleElement(2, 2), ShouldEqual, 3.0)
	})

	Convey("Iterator should yield values correctly.", t, func() {
		a := NewCSRMatrix(3, 3)
		a.Insert(0, 1, 1.0)
		a.Insert(2, 2, 3.0)
		a.Insert(1, 0, 2.0)
		iter := a.IterTriplets()

		triplet, ok := iter.Next()
		So(ok, ShouldEqual, true)
		So(triplet.row, ShouldEqual, 0)
		So(triplet.col, ShouldEqual, 1)
		So(triplet.val, ShouldEqual, 1.0)

		triplet, ok = iter.Next()
		fmt.Println(triplet)
		So(ok, ShouldEqual, true)
		So(triplet.row, ShouldEqual, 2)
		So(triplet.col, ShouldEqual, 2)
		So(triplet.val, ShouldEqual, 3.0)

		triplet, ok = iter.Next()
		So(ok, ShouldEqual, true)
		So(triplet, ShouldEqual, nil)
	})
}
