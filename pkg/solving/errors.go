package solving

import (
	"fmt"
)

type SolverError struct {
	Problem string
	Err error
}

func (e *SolverError) Unwrap() error {
	return e.Err
}

func (e *SolverError) Error() string {
	return fmt.Sprintf("Problem: %v, Error: %+v\n", e.Problem, e.Err)
}
