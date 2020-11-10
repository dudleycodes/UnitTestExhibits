package product

import (
	"errors"
	"mocking/postgresql"
)

var (
	// ErrNoSKUs is a sentinel when there are currently no available products.
	ErrNoSKUs = errors.New("there are no SKUs available at this time")
)

// ProductService is for interacting with the supplier's database.
type ProductService struct {
	database postgresql.PostgreSQL
}

// New creates a Product Service using the provided database connection.
func New(db postgresql.PostgreSQL) ProductService {
	return ProductService{
		database: db,
	}
}

// CountWidgetSKUs in the supplier's database; modifying the results to account for all available colors.
func (ps ProductService) CountWidgetSKUs() (int, error) {
	c, err := ps.database.RowCount("widgets")

	if err != nil {
		return 0, err
	}

	// If no  Stock Keep Units (SKUs) found, return the sentinel error that there are no products available.
	if c < 1 {
		return 0, ErrNoSKUs
	}

	// All the widgets come in three colors; but the supplier's database doesn't consider them different SKUs.
	return (c * 3), nil
}
