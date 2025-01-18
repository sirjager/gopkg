package utils

func Paginate[T any](items []T, limit, page int) []T {
	if limit == 0 {
		return items
	}
	start, end := PageEdges(len(items), limit, page)
	return items[start:end]
}

func PageEdges(total, limit, page int) (start, end int) {
	if limit == 0 {
		return 0, total
	}
	if page < 1 {
		page = 1
	}
	if limit > total {
		limit = total
	}

	start = (page - 1) * limit
	end = start + limit

	if start >= total {
		// no items
		return total, total
	}
	if end > total {
		end = total
	}
	return
}
