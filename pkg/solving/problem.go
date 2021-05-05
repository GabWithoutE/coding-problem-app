package solving

type Problem interface {
	Solve() (interface{}, error)
}
