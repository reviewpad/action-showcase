package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	tests := map[string]struct {
		args    []int
		wantVal []int
	}{
		"basic test": {
			args:    []int{5, 6, 7, 2, 1},
			wantVal: []int{1, 2, 5, 6, 7},
		},
		"2 test": {
			args:    []int{6, 5},
			wantVal: []int{5, 6},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gotVal := Quicksort(test.args)

			assert.Equal(t, test.wantVal, gotVal)
		})
	}
}
