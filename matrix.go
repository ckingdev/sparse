package sparse

// Matrix represents a generic type of matrix. Currently, only DOKMatrix and
// CSRMatrix implement this interface. In the future, methods will be
// provided that use Matrix parameters.
type Matrix interface {
	// Shape returns two ints describing the mxn shape of the matrix.
	Shape() (int, int)

	// Get returns the value in the matrix at the given indices.
	Get(i, j int) float64

	// Set inserts a new value or updates an old one at the given indices.
	Set(i, j int, val float64)

	// NNZ gives the number of nonzero entries in the matrix.
	NNZ() int
}
