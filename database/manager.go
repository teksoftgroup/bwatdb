package database

var initialPage uint64 = 0

// structure describing a page manager
type PageManager struct {
	MaxPage   uint64
	FreePages []uint64
}

func NewPageManager() *PageManager {
	return &PageManager{
		MaxPage:   initialPage,
		FreePages: []uint64{},
	}
}

func (pm *PageManager) GetNextPage() uint64 {
	if len(pm.FreePages) != 0 {
		pageID := pm.FreePages[len(pm.FreePages)-1]
		pm.FreePages = pm.FreePages[:len(pm.FreePages)-1]
		return pageID
	}
	pm.MaxPage += 1
	return pm.MaxPage
}

func (pm *PageManager) ReleasePage(pageNumber uint64) {
	pm.FreePages = append(pm.FreePages, pageNumber)
}
