package manager

import (
	"math"

	"github.com/fikrirnurhidayat/kasusastran/api"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
)

type PaginationManager interface {
	FromIncomingRequest(*api.RequestPagination) *Pagination
	NewOutgoingResponse(*Pagination) *api.ResponsePagination
}

type PaginationManagerImpl struct {
	defaultPageSize uint32
	defaultPage     uint32
}

func (pm *PaginationManagerImpl) FromIncomingRequest(req *api.RequestPagination) *Pagination {
	page := req.GetPage()
	pageSize := req.GetPageSize()

	if page == 0 {
		page = pm.defaultPage
	}

	if pageSize == 0 {
		pageSize = pm.defaultPageSize
	}

	res := &Pagination{
		Page:     page,
		PageSize: pageSize,
	}

	return res
}

func (pm *PaginationManagerImpl) NewOutgoingResponse(p *Pagination) (res *api.ResponsePagination) {
	page := p.Page
	pageSize := p.PageSize

	if page == 0 {
		page = pm.defaultPage
	}

	if pageSize == 0 {
		page = pm.defaultPageSize
	}

	return &api.ResponsePagination{
		Page:      page,
		PageSize:  pageSize,
		PageCount: p.PageCount,
		Total:     p.Total,
	}
}

type PaginationManagerSetter func(*PaginationManagerImpl)

func NewPaginationManager(setters ...PaginationManagerSetter) PaginationManager {
	pm := new(PaginationManagerImpl)

	for _, set := range setters {
		set(pm)
	}

	return pm
}

func WithDefaultPageSize(defaultPageSize uint32) PaginationManagerSetter {
	return func(pm *PaginationManagerImpl) {
		pm.defaultPageSize = defaultPageSize
	}
}

func WithDefaultPage(defaultPage uint32) PaginationManagerSetter {
	return func(pm *PaginationManagerImpl) {
		pm.defaultPage = defaultPage
	}
}

type Pagination struct {
	Page      uint32
	PageSize  uint32
	PageCount uint32
	Total     uint32
}

func (p *Pagination) WithTotal(total uint32) *Pagination {
	p.Total = total
	p.PageCount = uint32(math.Ceil(float64(p.Total) / float64(p.PageSize)))
	return p
}

func (p *Pagination) ToListQuery() *repository.ListQuery {
	limit := p.PageSize
	offset := (p.Page - 1) * limit

	return &repository.ListQuery{
		Limit:  limit,
		Offset: offset,
	}
}
