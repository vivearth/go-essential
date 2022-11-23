package challenge

import (
	"fmt"
	"math"
)

func evenendcount() int {
	count := 0

	for i := 1000; i < 10000; i++ {
		for j := i; j < 10000; j++ {
			prod := i * j
			// prodstr := fmt.Sprintf("%d", prod)
			// if prodstr[0] == prodstr[len(prodstr)-1] {
			// 	count++
			// }
			n := len(fmt.Sprintf("%d", prod)) - 1
			tempdiv := int(math.Pow(10, float64(n)))
			if (prod % 10) == (prod / tempdiv) {
				count++
			}
		}
	}
	return count
}
