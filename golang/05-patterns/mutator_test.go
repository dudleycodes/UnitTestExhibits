package patterns

import (
	"testing"
)

// A Mutator-Table Driven test contains a table of mutators, each which mutates a single field of a struct. Useful for
// targeting functions that validate structs.

func Test_OrderRequest_Validate(t *testing.T) {
	t.Parallel()

	// mutators contains a collection of functions that mutate a valid OrderRequest into an invalid one. Each
	// mutator only invalidates a single-field, so that all code paths inside the target's Validate() function
	// are covered.
	mutators := map[string]func(s *OrderRequest){
		"empty Address": func(s *OrderRequest) {
			s.Address = ShippingAddress{}
		},
		"empty Address FullName": func(s *OrderRequest) {
			s.Address.FullName = ""
		},
		"whitespace-only Address Address1": func(s *OrderRequest) {
			s.Address.Address1 = "   \r\n\t"
		},
		"no line items": func(s *OrderRequest) {
			s.LineItems = nil
		},
		"one line item with, a missing SKU": func(s *OrderRequest) {
			s.LineItems = []LineItem{
				LineItem{SKU: "", Quantity: 4},
			}
		},
		"two line items, one with a missing SKU": func(s *OrderRequest) {
			s.LineItems = []LineItem{
				LineItem{SKU: "EYE-PHONE-2-XL", Quantity: 1},
				LineItem{SKU: "", Quantity: 1},
			}
		},
		// Normally there would be many more mutators to cover all possible conditions where a OrderRequest
		// should fail validation.
	}

	happyPathChecked := false

	for name, mutator := range mutators {
		// The unmutated system under test (SUT) should always be valid, with a new one created for every
		// mutator to ensure no copy and/or reference errors skew the results.
		sut := &OrderRequest{
			Address: ShippingAddress{
				FullName: "H. J. Farnswoth",
				Address1: "57th Street",
				Address2: "The Angry Dome",
			},
			LineItems: []LineItem{
				LineItem{SKU: "POPPLER-BRE", Quantity: 92},
				LineItem{SKU: "L-UNIT", Quantity: 1},
			},
		}

		// Validate the unmutated sut to ensure Validate() recognizes it as error-free.
		if !happyPathChecked {
			if err := sut.Validate(); err != nil {
				t.Fatalf("A valid target should have passed but got error %q.", err.Error())
			}

			happyPathChecked = true
		}

		// Run through the mutators, ensuring each one causes Validate() to return an error.
		t.Run(name, func(t *testing.T) {
			mutator(sut)

			if err := sut.Validate(); err == nil {
				t.Error("Validation should have failed with an error, but got `nil`.")
			}
		})
	}
}
