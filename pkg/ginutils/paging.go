package ginutils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination contains information for pagination.
type Pagination struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

// Sort contains information for sorting.
type Sort struct {
	Field string `form:"sort_field"`
	Order string `form:"sort_order"` // asc or desc
}

// Search contains information for search criteria.
type Search struct {
	Query string `form:"search_query"`
}

func ParsePaginationSortSearch(c *gin.Context) (Pagination, Sort, Search) {
	var p Pagination
	var s Sort
	var search Search

	// Set default values
	p.Page = 1
	p.Limit = 10 // Default limit

	// Bind query parameters
	if pageStr := c.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			p.Page = page
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			p.Limit = limit
		}
	}

	s.Field = c.DefaultQuery("sort_field", "")
	s.Order = c.DefaultQuery("sort_order", "asc")
	search.Query = c.DefaultQuery("search_query", "")

	return p, s, search
}
