package sync

import (
	"errors"
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestArrayList_Add(t *testing.T) {
	testCases := []struct {
		name      string
		list      *ArrayList[int]
		index     int
		newVal    int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "Add num to index left",
			list:      NewArrayListOf([]int{1, 2, 3}),
			index:     0,
			newVal:    100,
			wantSlice: []int{100, 1, 2, 3},
		},
		{
			name:      "Add num to index right",
			list:      NewArrayListOf([]int{1, 2, 3}),
			index:     3,
			newVal:    100,
			wantSlice: []int{1, 2, 3, 100},
		},
		{
			name:      "Add num to index middle",
			list:      NewArrayListOf([]int{1, 2, 3}),
			index:     1,
			newVal:    100,
			wantSlice: []int{1, 100, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Add(tc.index, tc.newVal)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.list.vals, tc.wantSlice)
		})
	}
}

func TestArrayList_Cap(t *testing.T) {
	testCases := []struct {
		name    string
		wantRes int
		list    *ArrayList[int]
	}{
		{
			name:    "Cap empty list",
			wantRes: 5,
			list:    NewArrayListOf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cap := tc.list.Cap()

			assert.Equal(t, tc.wantRes, cap)
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	var a int
	fmt.Println(a)
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		index   int
		wantRes int
		wantErr error
	}{
		{
			name:    "Get index error",
			list:    NewArrayListOf([]int{1, 2, 3, 4}),
			index:   -1,
			wantRes: 0,
			wantErr: errors.New("index out of bounds"),
		},
		{
			name:    "Get index",
			list:    NewArrayListOf([]int{1, 2, 3, 4}),
			index:   0,
			wantRes: 1,
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.list.Get(tc.index)

			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
