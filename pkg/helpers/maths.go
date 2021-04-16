package helpers

/*
Minimum
params:
	x array of integers from which to find minimum
returns:
	int minimum value in the array
	error empty input
*/

func Minimum(x []int) int {
	m := x[0]
	for _, v := range x {
		if v < m {
			m = v
		}
	}

	return m
}