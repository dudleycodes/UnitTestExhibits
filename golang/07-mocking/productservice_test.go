package product

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"mocking/postgresql/mockgres"
)

func Test_CountWidgetSKUs(t *testing.T) {
	t.Parallel()

	t.Run("database error should bubble up", func(T *testing.T) {
		// the error expected to bubble up
		expected := fmt.Errorf("expected error %d", time.Now().UTC().Unix())

		// the postgresql mocker, with RowCount() behavior set to return the error
		mockDB := mockgres.New(mockgres.RowCountError(expected))

		// a real instance of product service, with the postgresql mocker injected
		sut := New(mockDB)

		if _, err := sut.CountWidgetSKUs(); !errors.Is(err, expected) {
			t.Fatalf("Expected error %q but got %q", expected, err)
		}
	})

	t.Run("No available SKUs should result in sentinel error `ErrNoSKUs`", func(T *testing.T) {
		// the postgresql mocker, with RowCount() behavior set to `0`.
		mockDB := mockgres.New(mockgres.RowCountValue(0))

		// a real instance of product service, with the postgresql mocker injected
		sut := New(mockDB)

		if _, err := sut.CountWidgetSKUs(); !errors.Is(err, ErrNoSKUs) {
			t.Fatalf("Expected error %q but got %q", ErrNoSKUs, err)
		}
	})

	t.Run("Count should get multiplied by three to account for all colors", func(T *testing.T) {
		tests := map[string]struct {
			expected int
			behavior mockgres.Behavior
		}{
			"3 should become 9": {
				expected: 9,
				behavior: mockgres.RowCountValue(3),
			},
			"12 should become 36": {
				expected: 36,
				behavior: mockgres.RowCountValue(12),
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				sut := New(mockgres.New(test.behavior))

				c, err := sut.CountWidgetSKUs()

				if err != nil {
					t.Fatalf("Got unexpected error: %s", err.Error())
				}

				if c != test.expected {
					t.Errorf("Expected the count to be `%d` but got `%d`", test.expected, c)
				}
			})
		}
	})
}

// Example of writing a custom behavior that's not provided with the mockgres package.
func Benchmark_CountWidgetSKUs(b *testing.B) {
	delayBehavior := func(b *mockgres.MockBehaviors) {
		b.RowCount = func(tableName string) (int, error) {
			if strings.HasPrefix(tableName, "view_") {
				time.Sleep(250 * time.Millisecond)
			} else {
				time.Sleep(20 * time.Millisecond)
			}

			return 30, nil
		}
	}

	sut := New(mockgres.New(delayBehavior))

	for n := 0; n < b.N; n++ {
		sut.CountWidgetSKUs()
	}
}
