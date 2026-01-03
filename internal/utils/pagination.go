package utils

type Pagination struct {
	Page       int
	Limit      int
	Offset     int
	TotalRows  int64
	TotalPages int
}

// function for add pagination
func NewPagination(page, limit int, totalRows int64) Pagination {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	totalPages := int((totalRows + int64(limit) - 1) / int64(limit))

	return Pagination{
		Page:       page,
		Limit:      limit,
		Offset:     (page - 1) * limit,
		TotalRows:  totalRows,
		TotalPages: totalPages,
	}
}
