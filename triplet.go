package sparse

// Triplet represents an element of a matrix- row, column, and value.
type Triplet struct {
	Row, Col int
	Val      float64
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

func (t *Triplet) LessThanIndices(i, j int) bool {
	if t.Row < i {
		return true
	}
	if t.Row == i && t.Col < j {
		return true
	}
	return false
}

func (t *Triplet) EqualIndices(i, j int) bool {
	if t.Row == i && t.Col == j {
		return true
	}
	return false
}
