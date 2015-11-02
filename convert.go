package sparse

func ConvertToCSC(m Matrix) *CompressedMatrix {
	switch a := m.(type) {
	case *CompressedMatrix:
		if a.isCSC {
			return a
		} else {
			// convert from CSR to CSC
		}
	case *DOKMatrix:
		return dokToCSC(a)
	case *COOMatrix:
		return cooToCSC(a)
	default:
		return nil
	}
	return nil
}

func dokToCSC(a *DOKMatrix) *CompressedMatrix {
	b := NewCSCMatrix(a.shape[0], a.shape[1])
	for rowInd, row := range a.data {
		if _, ok := a.data[rowInd]; !ok {
			a.data[rowInd] = map[int]float64{}
		}
		for colInd, val := range row {
			a.data[rowInd][colInd] = val
		}
	}
	return b
}

func dokToCSR(a *DOKMatrix) *CompressedMatrix {
	b := NewCSRMatrix(a.shape[0], a.shape[1])
	for rowInd, row := range a.data {
		for colInd, val := range row {
			b.Set(rowInd, colInd, val)
		}
	}
	return b
}

func cooToCSC(a *COOMatrix) *CompressedMatrix {
	b := NewCSCMatrix(a.shape[0], a.shape[1])
	for _, tri := range a.elements {
		b.Set(tri.Row, tri.Col, tri.Val)
	}
	return b
}

func csrToCSC(a *CompressedMatrix) *CompressedMatrix {
	b := NewCSCMatrix(a.shape[0], a.shape[1])
	iter := a.IterTriplets()
	for t, ok := iter.Next(); ok; t, ok = iter.Next() {
		b.Set(t.Row, t.Col, t.Val)
	}
	return b
}
