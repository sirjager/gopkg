package storage

import (
	"context"
	"errors"
	"reflect"
	"sort"
	"sync"

	"github.com/sirjager/gopkg/utils"
)

type _item struct {
	Value interface{}
}

type _storage struct {
	mu   sync.RWMutex
	data map[string]*_item
}

// NewInMemoryStorage creates a new instance of InMemoryStorage.
func NewInMemoryStorage() Storage {
	return &_storage{
		data: make(map[string]*_item),
	}
}

// Create stores a value at the given id.
func (s *_storage) Create(ctx context.Context, id string, value interface{}) (_ error) {
	_, exists := s.get(id)
	if exists {
		return errors.New(ErrIDAlreadyExits)
	}
	s.set(id, value)
	return nil
}

// List retrieves all values into values
func (s *_storage) List(ctx context.Context, values interface{}, limit, page int) (_ error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	rv := reflect.ValueOf(values)
	if rv.Kind() != reflect.Pointer || rv.IsNil() || rv.Elem().Kind() != reflect.Slice {
		return errors.New(ErrValueMustBeAPointerToSlice)
	}

	// extracting keys and sorting to maintain order
	keys := make([]string, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	keysToGet := utils.Paginate(keys, limit, page)

	slice := rv.Elem()
	// Populate the slice with corresponding items
	for _, key := range keysToGet {
		if item, exists := s.data[key]; exists {
			slice = reflect.Append(slice, reflect.ValueOf(item.Value))
		}
	}

	rv.Elem().Set(slice)
	return nil
}

// Read retrieves a value for a given id into value interface
func (s *_storage) Read(ctx context.Context, id string, value interface{}) (_ error) {
	item, exists := s.get(id)
	if !exists {
		return errors.New(ErrNotFound)
	}
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return errors.New(ErrValueMustBeAPointer)
	}
	rv.Elem().Set(reflect.ValueOf(item.Value))
	return nil
}

// Update updates an existing value at the given id.
func (s *_storage) Update(ctx context.Context, id string, value interface{}) (_ error) {
	_, exists := s.get(id)
	if !exists {
		return errors.New(ErrNotFound)
	}
	s.set(id, value)
	return nil
}

// Delete removes a value at the given id.
func (s *_storage) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, id)
	return nil
}

// Reset removes a value at the given id.
func (s *_storage) Reset(ctx context.Context) (_ error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[string]*_item)
	return nil
}

func (s *_storage) Count(ctx context.Context) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	count := len(s.data)
	return count, nil
}

func (s *_storage) get(id string) (*_item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, exists := s.data[id]
	return item, exists
}

func (s *_storage) set(id string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[id] = &_item{Value: value}
}
