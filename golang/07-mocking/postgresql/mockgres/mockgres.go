// Package mockgres provides a mock that is 1 to 1 compatible with the postgresql library, for use in unit tests.
package mockgres

// Behavior configures how a mock postgresql function should behave.
type Behavior func(b *MockBehaviors)

// MockBehaviors determines how all of Mocker's mock postgresql functions should behave.
type MockBehaviors struct {
	RowCount func(string) (int, error)
}

// New creates a Mocker configured with any user-provided behaviors for its mock functions. By default all functions
// will return non-error results.
func New(bx ...Behavior) Mocker {
	behaviors := MockBehaviors{}

	for _, b := range bx {
		if b != nil {
			b(&behaviors)
		}
	}

	return Mocker{
		behaviors: &behaviors,
	}
}

// Mocker is a mock PostgreSQL implementation that is 1 to 1 compatible with the postgresql.PostgreSQL interface.
type Mocker struct {
	behaviors *MockBehaviors
}

// Close does nothing in this example; it's defined to keep things 1 to 1 compatible with the postgresql.PostgreSQL
// interface.
func (m Mocker) Close() error {
	return nil
}

// RowCount mocks the postgresql library's `RowCount` function, will return (10, nil) if no behavior is configured.
func (m Mocker) RowCount(tableName string) (int, error) {
	if m.behaviors.RowCount != nil {
		return m.behaviors.RowCount(tableName)
	}

	return 10, nil
}

// RowCountError sets up the mock RowCount() function to return the provided error.
func RowCountError(e error) Behavior {
	return func(b *MockBehaviors) {
		b.RowCount = func(string) (int, error) {
			return -1, e
		}
	}
}

// RowCountValue sets up the mock RowCount() function to return a specific value and a nil error.
func RowCountValue(v int) Behavior {
	return func(b *MockBehaviors) {
		b.RowCount = func(string) (int, error) {
			return v, nil
		}
	}
}

// More customization examples

// With will modify the behaviors on an existing Mocker instance.
func (m *Mocker) With(bx ...Behavior) {
	for _, b := range bx {
		if b != nil {
			b(m.behaviors)
		}
	}
}

// RowCount injects a function that establishes the results the mock RowCount() function.
func RowCount(f func(string) (int, error)) Behavior {
	return func(b *MockBehaviors) {
		b.RowCount = f
	}
}
