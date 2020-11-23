package pagination

const PaginationCtxKey = "pagination"

type Pagination struct {
	Offset int
	Count  int
	Total  int
	Data   []interface{}
}

type PaginationContext struct {
	Limit  int
	Offset int
}
