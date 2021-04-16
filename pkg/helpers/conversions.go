package helpers

import (
	"github.com/pkg/errors"
	"strconv"
)

func Atoiarray (ss []string) ([]int, error) {
	var is []int

	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.Wrap(err, "Atoiarray: strconv failed")
		}
		is = append(is, i)
	}

	return is, nil
}