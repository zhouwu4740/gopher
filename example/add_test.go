package example

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	data := map[string][]int{
		"1+2=3": {1, 2, 3},
		"2+2=4": {2, 2, 4},
		"3+3=6": {3, 3, 6},
	}

	for name, tc := range data {
		t.Run(name, func(t *testing.T) {
			c := Add(tc[0], tc[1])
			assert.Equal(t, tc[2], c)
		})
	}
}
