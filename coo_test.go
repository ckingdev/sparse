package sparse

import (
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCOOMatrix(t *testing.T) {
	Convey("Inserted values should be retrieved correctly.", t, func() {
		a := NewCOOMatrix(3, 3)
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
}

func BenchmarkCOOInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := NewCOOMatrix(50, 50)
		vals := make([]Triplet, 250)
		for i := 0; i < 250; i++ {
			vals[i].Col = rand.Int() % 50
			vals[i].Row = rand.Int() % 50
			vals[i].Val = rand.Float64()
		}
		b.StartTimer()
		for _, t := range vals {
			a.Set(t.Row, t.Col, t.Val)
		}
	}
}
