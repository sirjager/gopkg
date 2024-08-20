package utils

// Generic function to check if a value exists in a slice
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

func GetFirstOrFallback[T any](fallback T, arr ...T) T {
	if len(arr) != 0 {
		return arr[0]
	}
	return fallback
}
