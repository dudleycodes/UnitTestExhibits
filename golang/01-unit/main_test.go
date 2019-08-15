package main

import (
	"testing"
)

// This test does its job, but when it fails we don't get any information about why or how.
func Test_01_Bad_Sum(t *testing.T) {
	if Sum(1, 2, 3, 4) != 10 || Sum(1, 2, 3, 4, -5) != 5 || Sum(2, -5) != -3 {
		t.Fail()
	}
}

// This test works and gives us some information about what failed, but the code is tedious not only
// for adding/editing/removing test cases, but also to read.
func Test_02_Tedious_Sum(t *testing.T) {
	// test case 1 -- basic addition
	actual, expected := Sum(1, 2, 3, 4), 10
	if actual != expected {
		t.Errorf("Should have resulted in the value of %v not %v", expected, actual)
	}

	//  test case 2 - addition with negative numbers
	actual, expected = Sum(1, 2, 3, 4, -5), 5
	if actual != expected {
		t.Errorf("Should have resulted in the value of %v not %v", expected, actual)
	}

	//  test case 3 - negative result
	actual, expected = Sum(2, -5), -3
	if actual != expected {
		t.Errorf("Should have resulted in the value of %v not %v", expected, actual)
	}
}

// This test works and cuts out a lot of the tedium but:
//	• it's hard to add/edit/remove test cases because it mixes the test-data with logic. When
//		writing tests, we only want to focus on the testing data (inputs/outputs).
//	• the test results still don't show us what inputs break things (only shows expected/actual)
func Test_03_LessTedious_Sum(t *testing.T) {
	testAreEqual := func(actual, expected int) {
		if actual != expected {
			t.Errorf("Should have resulted in the value of %d not %d", expected, actual)
		}
	}

	testAreEqual(1, 1)                   // meta test 0 - verify test function works
	testAreEqual(Sum(1, 2, 3, 4), 10)    // test case 1 - basic addition
	testAreEqual(Sum(1, 2, 3, 4, -5), 5) // test case 2 - addition with negative numbers
	testAreEqual(Sum(2, -5), -3)         // test case 3 - negative result
}

// This "table test" works; by isolating the logic we can focus just on test data when adding/editing/removing.
// More importantly, we can show all data around test failures.
func Test_04_TableTest_Sum(t *testing.T) {
	tests := []struct {
		inputs   []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 10},    // basic addition
		{[]int{1, 2, 3, 4, -5}, 5}, // addition with negative numbers
		{[]int{2, -5}, -3},         // negative result
	}

	for _, test := range tests {
		actual := Sum(test.inputs...)
		if actual != test.expected {
			t.Errorf("Sum of %v should have resulted in the value of %d not %d", test.inputs, test.expected, actual)
		}
	}
}

// Named table tests have all the benefits of normal table tests. By providing each test case with a name, we
// can ensure the motivations behind the tests are also displayed in the test results (and not just code comments
// like the earlier examples).
func Test_05_NamedTableTest_Sum(t *testing.T) {
	tests := map[string]struct {
		inputs   []int
		expected int
	}{
		"basic addition":                 {[]int{1, 2, 3, 4}, 10},
		"addition with negative numbers": {[]int{1, 2, 3, 4, -5}, 5},
		"negative result":                {[]int{2, -5}, -3},
	}

	for caseName, test := range tests {
		t.Run(caseName, func(t *testing.T) {
			actual := Sum(test.inputs...)
			if actual != test.expected {
				t.Errorf("Sum of %v should have resulted in the value of %d not %d", test.inputs, test.expected, actual)
			}
		})
	}
}
