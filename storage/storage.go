package storage

import "context"

// Storage interface defines methods for interacting with a storage system.
type Storage interface {
	// Create stores a value at the given id.
	Create(ctx context.Context, id string, value interface{}) error

	// List retrieves all values into values, limit and page can be passed to get a paginated list. Both are default to 0
	List(ctx context.Context, values interface{}, limit, page int) error

	// Read retrieves a value for a given id into value interface
	Read(ctx context.Context, id string, value interface{}) error

	// Update updates an existing value at the given id.
	Update(ctx context.Context, id string, value interface{}) error

	// Delete removes a value at the given id.
	Delete(ctx context.Context, id string) error

	// Reset storage
	Reset(ctx context.Context) error

	// Count returns the number of stored values
	Count(ctx context.Context) (int, error)
}
