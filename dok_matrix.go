package sparse

import (
	"fmt"
)

type DOKMatrix struct {
	data  map[int]map[int]float64
	shape [2]int
}

func NewDOKMatrix(m, n int) *DOKMatrix {
	return &DOKMatrix{
		data:  make(map[int]map[int]float64),
		shape: [2]int{m, n},
	}
}

func (d *DOKMatrix) Get(i, j int) float64 {
	row, ok := d.data[i]
	if !ok {
		return 0.0
	}
	return row[j]
}

func (d *DOKMatrix) Shape() (int, int) {
	return d.shape[0], d.shape[1]
}

func (d *DOKMatrix) Set(i, j int, val float64) {
	row, ok := d.data[i]
	if !ok {
		d.data[i] = map[int]float64{j: val}
		return
	}
	row[j] = val
}

func (d *DOKMatrix) Copy() *DOKMatrix {
	mcopy := NewDOKMatrix(d.shape[0], d.shape[1])
	for i, row := range mcopy.data {
		mcopy.data[i] = map[int]float64{}
		for j, val := range row {
			mcopy.data[i][j] = val
		}
	}
	return mcopy
}

func DOKSum(m1, m2 *DOKMatrix) (*DOKMatrix, error) {
	if m1.shape[0] != m2.shape[0] || m1.shape[1] != m2.shape[1] {
		return nil, fmt.Errorf("Size mismatch")
	}
	sum := m1.Copy()
	for i, row := range m2.data {
		for j, val := range row {
			sumRow, ok := sum.data[i]
			if !ok {
				sum.data[i] = map[int]float64{j: val}
				continue
			}
			sumRow[j] += val
		}
	}
	return sum, nil
}

func (d *DOKMatrix) NNZ() int {
	nnz := 0
	for _, row := range d.data {
		nnz += len(row)
	}
	return nnz
}
