package sparse

type Matrix interface {
	Index(i, j int) float64
	Size() (int, int)
	Insert(i, j int, val float64)
	NNZ() int
}
