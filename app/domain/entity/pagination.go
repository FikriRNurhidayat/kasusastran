package entity

import "math"

type Pagination struct {
	Page      uint32
	PageSize  uint32
	PageCount uint32
	Total     uint32
}

func (e *Pagination) GetLimit() uint32 {
	return e.PageSize
}

func (e *Pagination) GetOffset() uint32 {
	var page uint32

	if e.Page > 0 {
		page = e.Page - 1
	}

	return page * e.PageSize
}

func NewPagination(limit uint32, offset uint32, count uint32) *Pagination {
	ilimit := float64(limit)
	ioffset := float64(offset)
	icount := float64(count)

	Page := math.Floor(ioffset/ilimit) + 1
	PageSize := ilimit
	PageCount := math.Ceil(icount / ilimit)

	return &Pagination{
		Page:      uint32(Page),
		PageSize:  uint32(PageSize),
		PageCount: uint32(PageCount),
		Total:     uint32(icount),
	}
}
