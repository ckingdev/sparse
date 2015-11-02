package sparse

type COOMatrix struct {
	elements []Triplet
	shape    [2]int
}

func (c *COOMatrix) Shape() (int, int) {
	return c.shape[0], c.shape[1]
}

func (c *COOMatrix) NNZ() int {
	return len(c.elements)
}

func (c *COOMatrix) Get(i, j int) float64 {
	for _, t := range c.elements {
		if t.LessThanIndices(i, j) {
			continue
		}
		if t.EqualIndices(i, j) {
			return t.Val
		}
		return 0.0
	}
	return 0.0
}

func (c *COOMatrix) Set(i, j int, val float64) {
	for ind, t := range c.elements {
		if t.LessThanIndices(i, j) {
			continue
		}
		if t.EqualIndices(i, j) {
			c.elements[ind].Val = val
			return
		}
		// need to insert, we're greater than the given indices
		c.elements = append(c.elements, Triplet{})
		for k := len(c.elements) - 2; k >= ind; k-- {
			c.elements[k+1] = c.elements[k]
		}
		c.elements[ind] = Triplet{
			Row: i,
			Col: j,
			Val: val,
		}
		return
	}
	c.elements = append(c.elements, Triplet{
		Row: i,
		Col: j,
		Val: val,
	})
}

func NewCOOMatrix(m, n int) *COOMatrix {
	return &COOMatrix{
		elements: []Triplet{},
		shape:    [2]int{m, n},
	}
}
