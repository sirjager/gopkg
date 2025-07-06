package utils

import (
	"reflect"
)

// IsEmpty checks if the given value is empty.
// It handles different types, including strings, slices, arrays, maps, channels, pointers, and interfaces.
// For strings, it checks if the string is empty.
// For slices, arrays, and maps, it checks if the length is zero.
// For channels, it checks if the length is zero.
// For pointers and interfaces, it checks if the value is nil.
// For all other types, it defaults to false, assuming they are not empty.
func IsEmpty[T any](value T) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Slice, reflect.Array:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	case reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

// ValueExist checks if a value exists in a slice
func ValueExist[T comparable](find T, in []T) bool {
	for _, v := range in {
		if v == find {
			return true
		}
	}
	return false
}

// The ReverseArray function takes a list and
// returns a new list with the elements in reverse order.
func ReverseArray[T any](list []T) []T {
	length := len(list)
	reversed := make([]T, length)
	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversed[i] = list[j]
	}
	return reversed
}

// RemoveDuplicates removes duplicates from a list
func RemoveDuplicates[T comparable](list []T) []T {
	uniqueMap := make(map[T]bool)
	uniqueList := []T{}
	for _, item := range list {
		if !uniqueMap[item] {
			uniqueMap[item] = true
			uniqueList = append(uniqueList, item)
		}
	}
	return uniqueList
}

// GetFirstOrFallback returns the first element of the array if it is not empty, otherwise returns the fallback value
func GetFirstOrFallback[T any](fallback T, arr ...T) T {
	if len(arr) != 0 {
		return arr[0]
	}
	return fallback
}
