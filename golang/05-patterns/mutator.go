package patterns

import (
	"errors"
	"strings"
)

// OrderRequest represents a customer's request for order, received via external integrations.
type OrderRequest struct {
	// Address where the order should be shipped to.
	Address ShippingAddress
	// LineItems contains the breakdown of the items to be ordered.
	LineItems []LineItem
}

// LineItem represents a line of product the customer wants to purchase, received via external integrations.
type LineItem struct {
	// Quantity indicates how many items should be ordered.
	Quantity int
	// SKU (Stock Keeping Unit) is the unique ID of a product that can be ordered.
	SKU string
}

// ShippingAddress represents the destinantion details for a delivery.
type ShippingAddress struct {
	// FullName that the delivery should be addressed to (required).
	FullName string
	// Address1 is the street or P.O. Box the delivery should be addressed to (required).
	Address1 string
	// Address2 is optional information about where a delivery should be address to (apartment, suite, floor, etc).
	Address2 string
}

// Validate an order request is acceptable before accepting it into the system.
func (req OrderRequest) Validate() error {
	if req.Address == (ShippingAddress{}) {
		return errors.New("A shipping address must be provided")
	}

	if strings.TrimSpace(req.Address.FullName) == "" {
		return errors.New("Shipping address must have a `FullName`")
	}

	if strings.TrimSpace(req.Address.Address1) == "" {
		return errors.New("Shipping address must have an `Address1`")
	}

	if len(req.LineItems) < 1 {
		return errors.New("An order must contain at least one `LineItem`")
	}

	for _, li := range req.LineItems {
		if li.Quantity < 1 {
			return errors.New("A `LineItem` must have a quantity of at least 1")
		}

		if strings.TrimSpace(li.SKU) == "" {
			return errors.New("A `LineItem` must include a `SKU`")
		}
	}

	return nil
}
