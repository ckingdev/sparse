package sparse

type Item struct {
	Row int
	Col int
	Val float64
}

type Matrix interface {
	Index(i, j int) float64
	Size() (int, int)
	Insert(i, j int, val float64)
	IterItems() chan Item
	IterValues() chan float64
	NNZ() int
}
