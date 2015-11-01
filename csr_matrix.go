package sparse

// CSRMatrix represents a sparse matrix in compressed row storage format.
type CSRMatrix struct {
	// data in row-major format
	data []float64

	// column indices of values in CSRMatrix.data
	indices []int

	// index in data of first element in each row
	indptr []int

	shape [2]int
}

// Shape returns two ints describing the mxn shape of the matrix.
func (c *CSRMatrix) Shape() (int, int) {
	return c.shape[0], c.shape[1]
}

// NewCSRMatrix creates a new CSRMatrix with the given shape.
func NewCSRMatrix(m, n int) *CSRMatrix {
	return &CSRMatrix{
		data:    []float64{},
		indices: []int{},
		indptr:  make([]int, m+1),
		shape:   [2]int{m, n},
	}
}

// Get returns the value in the matrix at the given indices.
func (c *CSRMatrix) Get(row, col int) float64 {
	majorIndex := row
	minorIndex := col
	start := c.indptr[majorIndex]
	end := c.indptr[majorIndex+1]
	for i := start; i < end; i++ {
		if c.indices[i] == minorIndex {
			return c.data[i]
		}
	}
	return 0.0
}

// Set inserts a new value or updates an old one at the given indices.
func (c *CSRMatrix) Set(row, col int, val float64) {
	begin := c.indptr[row]
	end := c.indptr[row+1]
	if begin == end {
		c.indices = append(c.indices[0:begin], append([]int{col}, c.indices[begin:]...)...)
		c.data = append(c.data[0:begin], append([]float64{val}, c.data[begin:]...)...)
		for i := row + 1; i < c.shape[0]; i++ {
			c.indptr[i]++
		}
		c.indptr[c.shape[0]]++
		return
	}
	for i := begin; i < end; i++ {
		if c.indices[i] == col {
			c.data[i] = val
			return
		} else if c.indices[i] > col {
			// not present- need to insert
			for j := i + 1; j < c.shape[0]; j++ {
				c.indptr[j]++
			}
			c.indices = append(c.indices[0:i], append([]int{col}, c.indices[i:]...)...)
			c.data = append(c.data[0:i], append([]float64{val}, c.data[i:]...)...)
			c.indptr[c.shape[0]]++
		}
	}
	c.indices = append(c.indices[0:end], append([]int{col}, c.indices[end:]...)...)
	c.data = append(c.data[0:end], append([]float64{val}, c.data[end:]...)...)
	for j := row + 1; j <= c.shape[0]; j++ {
		c.indptr[j]++
	}
}

// NNZ gives the number of nonzero entries in the matrix.
func (c *CSRMatrix) NNZ() int {
	return c.indptr[c.shape[0]]
}

// CSRIterator represents an iterator that yields the non-zero values in the
// matrix in row-major order.
type CSRIterator struct {
	m        *CSRMatrix
	valIndex int
	rowIndex int
	rowStart int
	rowEnd   int
}

// IterTriplets creates a new iterator that will yield the value of the matrix.
func (c *CSRMatrix) IterTriplets() *CSRIterator {
	return &CSRIterator{
		m:        c,
		valIndex: 0,
		rowIndex: 0,
		rowStart: 0,
		rowEnd:   c.indptr[1],
	}
}

// Triplet represents an element of a matrix- row, column, and value.
type Triplet struct {
	Row, Col int
	Val      float64
}

// Next yields the next element of the matrix (using row-major ordering) and
// advances the iterator.
func (t *CSRIterator) Next() (*Triplet, bool) {
	if t.valIndex >= t.m.indptr[t.m.shape[0]] {
		return nil, false
	}
	for t.rowStart == t.rowEnd {
		t.rowIndex++
		t.rowStart = t.m.indptr[t.rowIndex]
		t.rowEnd = t.m.indptr[t.rowIndex+1]
	}
	ret := &Triplet{
		Row: t.rowIndex,
		Col: t.m.indices[t.valIndex],
		Val: t.m.data[t.valIndex],
	}
	t.valIndex++
	if t.valIndex == t.rowEnd {
		for t.valIndex == t.rowEnd && t.rowEnd != t.m.indptr[t.m.shape[0]] {
			t.rowIndex++
			t.rowStart = t.m.indptr[t.rowIndex]
			t.rowEnd = t.m.indptr[t.rowIndex+1]
		}
	}
	return ret, true
}

// LessThan returns whether the calling triplet is less than another triplet
// given row-major ordering.
func (t *Triplet) LessThan(other *Triplet) bool {
	if t.Row < other.Row {
		return true
	}
	if t.Row == other.Row && t.Col < other.Col {
		return true
	}
	return false
}

// AddCSR computes the sum of two CSR matrices.
func AddCSR(c1 *CSRMatrix, c2 *CSRMatrix) *CSRMatrix {
	if c1.shape[0] != c2.shape[0] || c1.shape[1] != c2.shape[1] {
		panic("Adding matrices of different sizes")
	}
	var larger int
	if c1.NNZ() > c2.NNZ() {
		larger = c1.NNZ()
	} else {
		larger = c2.NNZ()
	}
	data := make([]float64, 0, larger)
	indptr := make([]int, c1.shape[0]+1)
	indices := make([]int, 0, larger)
	iter1 := c1.IterTriplets()
	iter2 := c2.IterTriplets()
	t1, ok1 := iter1.Next()
	t2, ok2 := iter2.Next()
	for {
		if !ok1 && !ok2 {
			break
		}
		if !ok1 {
			for ok2 {
				data = append(data, t2.Val)
				indices = append(indices, t2.Col)
				for i := t2.Row + 1; i < c2.shape[0]+1; i++ {
					indptr[i]++
				}
				t2, ok2 = iter2.Next()
			}
			break
		}
		if !ok2 {
			for ok1 {
				data = append(data, t1.Val)
				indices = append(indices, t1.Col)
				for i := t1.Row + 1; i < c1.shape[0]+1; i++ {
					indptr[i]++
				}
				t1, ok1 = iter1.Next()
			}
			break
		}
		if t1.LessThan(t2) {
			// add t1, advance t1, and continue
			data = append(data, t1.Val)
			indices = append(indices, t1.Col)
			for i := t1.Row + 1; i < c1.shape[0]+1; i++ {
				indptr[i]++
			}
			t1, ok1 = iter1.Next()
			continue
		} else if t2.LessThan(t1) {
			// add t2, advance t2, and continue
			data = append(data, t2.Val)
			indices = append(indices, t2.Col)
			for i := t2.Row + 1; i < c2.shape[0]+1; i++ {
				indptr[i]++
			}
			t2, ok2 = iter2.Next()
			continue
		} else {
			// add t1+t2, advance both, continue
			data = append(data, t1.Val+t2.Val)
			indices = append(indices, t1.Col)
			for i := t1.Row + 1; i < c1.shape[0]+1; i++ {
				indptr[i]++
			}
			t1, ok1 = iter1.Next()
			t2, ok2 = iter2.Next()
			continue
		}
	}
	return &CSRMatrix{
		data:    data,
		indptr:  indptr,
		indices: indices,
		shape:   c1.shape,
	}
}
