package sparse

// DOKMatrix represents a sparse matrix in a dictionary of keys format. i.e.,
// a map of row -> column -> value.
type DOKMatrix struct {
	data  map[int]map[int]float64
	shape [2]int
}

// NewDOKMatrix creates a new DOKMatrix with the given shape.
func NewDOKMatrix(m, n int) *DOKMatrix {
	return &DOKMatrix{
		data:  make(map[int]map[int]float64),
		shape: [2]int{m, n},
	}
}

// Get returns the value in the matrix at the given indices.
func (d *DOKMatrix) Get(i, j int) float64 {
	row, ok := d.data[i]
	if !ok {
		return 0.0
	}
	return row[j]
}

// Shape returns two ints describing the mxn shape of the matrix.
func (d *DOKMatrix) Shape() (int, int) {
	return d.shape[0], d.shape[1]
}

// Set inserts a new value or updates an old one at the given indices.
func (d *DOKMatrix) Set(i, j int, val float64) {
	_, ok := d.data[i]
	if !ok {
		d.data[i] = map[int]float64{j: val}
		return
	}
	d.data[i][j] = val
}

// Copy creates a deep copy of the calling matrix.
func (d *DOKMatrix) Copy() *DOKMatrix {
	mcopy := NewDOKMatrix(d.shape[0], d.shape[1])
	for i, row := range d.data {
		mcopy.data[i] = map[int]float64{}
		for j, val := range row {
			mcopy.data[i][j] = val
		}
	}
	return mcopy
}

// AddDOK computes the sum of two DOK matrices.
func AddDOK(m1, m2 *DOKMatrix) *DOKMatrix {
	if m1.shape[0] != m2.shape[0] || m1.shape[1] != m2.shape[1] {
		panic("Size mismatch")
	}
	sum := m1.Copy()
	for i, row := range m2.data {
		_, ok := sum.data[i]
		if !ok {
			sum.data[i] = map[int]float64{}
		}
		for j, val1 := range row {
			val2, ok := sum.data[i][j]
			if !ok {
				sum.data[i][j] = val1
				continue
			}
			sum.data[i][j] = val1 + val2
		}
	}
	return sum
}

// NNZ gives the number of nonzero entries in the matrix.
func (d *DOKMatrix) NNZ() int {
	nnz := 0
	for _, row := range d.data {
		nnz += len(row)
	}
	return nnz
}
