package challenge

import (
	"fmt"
)

type Ordere interface {
	int | float64 | string
}

func Min[T Ordere](values []T) (T, error) {
	if len(values) == 0 {
		var zeroval T
		return zeroval, fmt.Errorf("min of empty slice")
	}
	mini := values[0]
	for _, val := range values[1:] {
		if val < mini {
			mini = val
		}
	}
	return mini, nil
}
