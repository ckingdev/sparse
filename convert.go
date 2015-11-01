package sparse

//func DOKtoCSR(d *DOKMatrix) *CSRMatrix {
//	c := NewCSRMatrix(d.shape[0], d.shape[1])
//	for i, row := range d.data {
//		for j, val := range row {
//			c.Set(i, j, val)
//		}
//	}
//	return c
//}

//func CSRtoDOK(c *CSRMatrix) *DOKMatrix {
//	d := NewDOKMatrix(c.shape[0], c.shape[1])
//	iter := c.IterTriplets()
//	for triplet, ok := iter.Next(); ok; triplet, ok = iter.Next() {
//		d.Set(triplet.Row, triplet.Col, triplet.Val)
//	}
//	return d
//}
