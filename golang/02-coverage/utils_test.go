package main

import (
	"strconv"
	"testing"
)

func Test_getPlaceValue(t *testing.T) {
	tests := []struct {
		in       int
		expected string
	}{
		{1, "ones"},
		{23, "tens"},
		{128, "hundreds"},
		{4250, "thousands"},
		{87178, "ten thousands"},
		{111111, "hundred thousands"},
		{12345667, "ten millions"},
		{987654321, "hundred millions"},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			actual := getPlaceValue(test.in)

			if actual != test.expected {
				t.Errorf("Expected %q but got %q.", test.expected, actual)
			}
		})
	}
}

func Test_reverseString(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"ABC123", "321CBA"},
		{"racecar", "racecar"},
		{"Hi-de-ho neighborino!", "!onirobhgien oh-ed-iH"},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			actual := reverseString(test.in)
			if actual != test.expected {
				t.Errorf("Expected %q but got %q.", test.expected, actual)
			}
		})
	}
}
