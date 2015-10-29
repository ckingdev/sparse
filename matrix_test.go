package sparse

import (
	"fmt"
	"testing"
)

func TestDOKMatrix(t *testing.T) {
	d := NewDOKMatrix(3, 3)
	d.Insert(2, 2, 1.5)
	if d.Index(2, 2) != 1.5 {
		t.FailNow()
	}
	if d.Index(0, 0) != 0.0 {
		t.FailNow()
	}
}

func TestCSRMatrix(t *testing.T) {
	c := NewCSRMatrix(3, 3)
	c.setSingleElement(0, 1, 1.0)
	c.setSingleElement(0, 1, 1.0)
	c.setSingleElement(0, 0, 2.0)
	c.setSingleElement(2, 2, 3.0)
	if c.getSingleElement(0, 0) != 2.0 {
		fmt.Printf("Expected c[0, 0] == 2.0, got %v\n", c.getSingleElement(0, 0))
		fmt.Printf("%v\n", c)
		t.FailNow()
	}
	if c.getSingleElement(0, 1) != 1.0 {
		fmt.Printf("Expected c[0, 1] == 1.0, got %v\n", c.getSingleElement(0, 0))
		fmt.Printf("%v\n", c)
		t.FailNow()
	}
}

func TestCSRMatrixAdd(t *testing.T) {
	c1 := NewCSRMatrix(3, 4)
	c2 := NewCSRMatrix(3, 4)
	c1.setSingleElement(0, 0, 1)
	c1.setSingleElement(1, 1, 1)
	c1.setSingleElement(2, 2, 1)
	c2.setSingleElement(0, 0, 1)
	c2.setSingleElement(0, 1, 1)
	c2.setSingleElement(0, 2, 1)
	fmt.Println(c1)
	fmt.Println(c2)
	c3 := AddCSR(c1, c2)
	fmt.Println(c3)
}
