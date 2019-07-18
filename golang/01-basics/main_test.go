package main

import (
	"testing"
)

//This test does its job but when it fails we don't get any information about why or how.
func Test_00_BadTest_addNumbers(t *testing.T) {
	if addNumbers(1, 2, 3, 4) != 10 || addNumbers(1, 2, 3, 4, -5) != 5 || addNumbers(2, -5) != -3 {
		t.Fail()
	}
}

//This test works and give us some information about what failed but the code is tedious to
//not only to add/edit/remove test cases, but also to read.
func Test_01_TediousTest_addNumbers(t *testing.T) {
	// test case 1 -- basic addition
	actual, expected := addNumbers(1, 2, 3, 4), 10
	if actual != expected {
		t.Errorf("Should have resulted in the value of %v not %v", expected, actual)
	}

	//  test case 2 - addition with negative numbers
	actual, expected = addNumbers(1, 2, 3, 4, -5), 5
	if actual != expected {
		t.Errorf("Should have resulted in the value of %v not %v", expected, actual)
	}

	//  test case 3 - negative result
	actual, expected = addNumbers(2, -5), -3
	if actual != expected {
		t.Errorf("Should have resulted in the value of %v not %v", expected, actual)
	}
}

//This test works and cuts out a lot of the tedium but:
//	• it's hard to add/edit/remove test cases because it mixes the test-data with logic. When writing
//	  tests we only want to focus on the testing data (inputs/outputs).
//	• the test results still don't show us what inputs break things (only shows expected/actual)
func Test_02_LessTediousTest_addNumbers(t *testing.T) {
	testAreEqual := func(actual, expected int) {
		if actual != expected {
			t.Errorf("Should have resulted in the value of %d not %d", expected, actual)
		}
	}

	testAreEqual(1, 1)                          // meta test 0 - verify test function works
	testAreEqual(addNumbers(1, 2, 3, 4), 10)    // test case 1 - basic addition
	testAreEqual(addNumbers(1, 2, 3, 4, -5), 5) // test case 2 - addition with negative numbers
	testAreEqual(addNumbers(2, -5), -3)         // test case 3 - negative result
}

//This "table test" works; by isolating the logic we can focus just on test data when add/editing/removing. More importantly
//we can show all data round test failures.
func Test_03_TableTest_addNumbers(t *testing.T) {
	tests := []struct {
		inputs   []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 10},    // basic addition
		{[]int{1, 2, 3, 4, -5}, 5}, // addition with negative numbers
		{[]int{2, -5}, -3},         // negative result
	}

	for _, test := range tests {
		actual := addNumbers(test.inputs...)
		if actual != test.expected {
			t.Errorf("Sum of %v should have resulted in the value of %d not %d", test.inputs, test.expected, actual)
		}
	}
}

//Named table tests have all the benefits of normal table tests. By providing each test case with a name we can ensure the motivations
//behind the tests are also displayed in the test results (and not just code comments like the earlier examples).
func Test_04_NamedTableTest_addNumbers(t *testing.T) {
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
			actual := addNumbers(test.inputs...)
			if actual != test.expected {
				t.Errorf("Sum of %v should have resulted in the value of %d not %d", test.inputs, test.expected, actual)
			}
		})
	}
}

//When developing/modifying high-milage functions (code that gets called a lot) it's smart to keep an eye on its performance.
func Benchmark_addNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addNumbers(1, 2, 3, 4, 10)
		addNumbers(1, 2, 3, 4, -5)
		addNumbers(2, -5)
	}
}
