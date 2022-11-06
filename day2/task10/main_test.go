package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type testCase struct {
		name     string
		x        int
		y        int
		expected int
	}

	testCases := []testCase{
		{name: "No1", x: 1, y: 1, expected: 2},
		{name: "No2", x: 2, y: 3, expected: 5},
		{name: "No3", x: 4, y: 6, expected: 10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			obj, _ := newObj()
			actual := obj.add(tc.x, tc.y)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

func TestSub(t *testing.T) {
	type testCase struct {
		name     string
		x        int
		y        int
		expected int
	}

	testCases := []testCase{
		{name: "No1", x: 1, y: 1, expected: 0},
		{name: "No2", x: 2, y: 3, expected: -1},
		{name: "No3", x: 4, y: 6, expected: -2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			obj, _ := newObj()
			actual := obj.sub(tc.x, tc.y)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
