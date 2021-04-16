package solving

type ProblemService interface {
	Solve() (interface{}, error)
}