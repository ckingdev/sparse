package sparse

import "fmt"

type CSRMatrix struct {
	// data in row-major format
	data []float64

	// column indices of values in CSRMatrix.data
	indices []int

	// index in data of first element in each row
	indptr []int

	shape [2]int
}

func (c *CSRMatrix) Index(i, j int) float64 {
	return c.getSingleElement(i, j)
}

func (c *CSRMatrix) Size() (int, int) {
	return c.shape[0], c.shape[1]
}

func (c *CSRMatrix) Insert(i, j int, val float64) {
	c.setSingleElement(i, j, val)
}

func (c *CSRMatrix) IterItems() chan Item {
	ret := make(chan Item)
	go func() {
		for i := 0; i < c.shape[0]; i++ {
			for j := c.indptr[i]; j < c.indptr[i+1]; j++ {
				ret <- Item{i, c.indices[j], c.data[j]}
			}
		}
		close(ret)
	}()
	return ret
}

func (c *CSRMatrix) IterValues() chan float64 {
	ret := make(chan float64)
	go func() {
		for _, val := range c.data {
			ret <- val
		}
		close(ret)
	}()
	return ret
}

func NewCSRMatrix(m, n int) *CSRMatrix {
	return &CSRMatrix{
		data:    []float64{},
		indices: []int{},
		indptr:  make([]int, m+1),
		shape:   [2]int{m, n},
	}
}

func (c *CSRMatrix) getSingleElement(row, col int) float64 {
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

func (c *CSRMatrix) setSingleElement(row, col int, val float64) {
	begin := c.indptr[row]
	end := c.indptr[row+1]
	if begin == end {
		c.indices = append(c.indices[0:begin], append([]int{col}, c.indices[begin:]...)...)
		c.data = append(c.data[0:begin], append([]float64{val}, c.data[begin:]...)...)
		for i := begin + 1; i < c.shape[0]; i++ {
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
}

func (c *CSRMatrix) NNZ() int {
	return c.indptr[c.shape[0]]
}

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
	for row, start1 := range c1.indptr[0 : len(c1.indptr)-1] {
		fmt.Println(row)
		indptr[row+1] = indptr[row]
		start2 := c1.indptr[row]
		end1 := c1.indptr[row+1]
		end2 := c2.indptr[row+1]
		if start1 == end1 {
			if start2 == end2 {
				continue
			}
			for k := start2; k < end2; k++ {
				data = append(data, c2.data[k])
				indices = append(indices, c2.indices[k])
				indptr[row+1]++
			}
			continue
		} else if start2 == end2 {
			for k := start1; k < end1; k++ {
				data = append(data, c1.data[k])
				indices = append(indices, c1.indices[k])
				indptr[row+1]++
			}
			continue
		}
		i := start1
		j := start2
		for {
			if i == end1 && j == end2 {
				break
			} else if i == end1 {
				for k := j; k < end2; k++ {
					data = append(data, c2.data[k])
					indices = append(indices, c2.indices[k])
					indptr[row+1]++
				}
				break
			} else if j == end2 {
				for k := i; k < end1; k++ {
					data = append(data, c1.data[k])
					indices = append(indices, c1.indices[k])
					indptr[row+1]++
				}
				break
			}
			if c1.indices[i] == c1.indices[j] {
				fmt.Printf("i, j: %v, %v\n", i, j)
				val := c1.data[i] + c2.data[j]
				data = append(data, val)
				indices = append(indices, c1.indices[i])
				i++
				j++
			} else if c1.indices[i] < c2.indices[j] {
				data = append(data, c1.data[i])
				indices = append(indices, c1.indices[i])
				indptr[row+1]++
				i++
			} else {
				data = append(data, c2.data[j])
				indices = append(indices, c2.indices[j])
				indptr[row+1]++
				j++
			}
			indptr[row+1]++
		}
	}
	return &CSRMatrix{
		data:    data,
		indptr:  indptr,
		indices: indices,
		shape:   c1.shape,
	}
}
