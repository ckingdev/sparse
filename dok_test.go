package sparse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDOKSet(t *testing.T) {
	Convey("DOK matrix should get and set correctly.", t, func() {
		a := NewDOKMatrix(2, 2)

		a.Set(0, 0, 1.0)
		So(a.Get(0, 0), ShouldEqual, 1.0)
		So(a.Get(1, 1), ShouldEqual, 0.0)

		a.Set(0, 1, 2.0)
		a.Set(1, 0, 3.0)
		So(a.Get(0, 0), ShouldEqual, 1.0)
		So(a.Get(0, 1), ShouldEqual, 2.0)
		So(a.Get(1, 0), ShouldEqual, 3.0)
		So(a.Get(1, 1), ShouldEqual, 0.0)

		a.Set(0, 0, 6.0)
		So(a.Get(0, 0), ShouldEqual, 6.0)
		So(a.Get(0, 1), ShouldEqual, 2.0)
		So(a.Get(1, 0), ShouldEqual, 3.0)
		So(a.Get(1, 1), ShouldEqual, 0.0)
	})
}

func TestAddDOK(t *testing.T) {
	Convey("Addition of two zero matrices should be a zero matrix.", t, func() {
		a := NewDOKMatrix(2, 2)
		b := NewDOKMatrix(2, 2)
		c := AddDOK(a, b)
		So(c.Get(0, 0), ShouldEqual, 0)
		So(c.Get(0, 1), ShouldEqual, 0)
		So(c.Get(1, 0), ShouldEqual, 0)
		So(c.Get(1, 1), ShouldEqual, 0)
	})
	Convey("Addition of two matrices with no common elements should be correct.", t, func() {
		a := NewDOKMatrix(2, 2)
		a.Set(0, 0, 1.0)
		a.Set(1, 1, 2.0)
		b := NewDOKMatrix(2, 2)
		b.Set(1, 0, 3.0)
		b.Set(0, 1, 4.0)
		c := AddDOK(a, b)
		t.Logf("%v", c)
		So(c.Get(0, 0), ShouldEqual, 1.0)
		So(c.Get(0, 1), ShouldEqual, 4.0)
		So(c.Get(1, 0), ShouldEqual, 3.0)
		So(c.Get(1, 1), ShouldEqual, 2.0)
	})
	Convey("Addition of two matrices with common elements should be correct.", t, func() {
		a := NewDOKMatrix(2, 2)
		a.Set(0, 0, 1.0)
		a.Set(1, 1, 2.0)
		a.Set(0, 1, 2.0)
		b := NewDOKMatrix(2, 2)
		b.Set(1, 0, 3.0)
		b.Set(0, 1, 4.0)
		b.Set(1, 1, 1.0)
		c := AddDOK(a, b)
		So(c.Get(0, 0), ShouldEqual, 1.0)
		So(c.Get(0, 1), ShouldEqual, 6.0)
		So(c.Get(1, 0), ShouldEqual, 3.0)
		So(c.Get(1, 1), ShouldEqual, 3.0)
	})
}
