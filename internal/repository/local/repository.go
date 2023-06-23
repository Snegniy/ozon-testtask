package local

import (
	"fmt"
	"github.com/Snegniy/ozon-testtask/internal/model"
	"sync"
)

type Repository struct {
	storageBase  map[string]model.UrlBaseStorage
	storageShort map[string]model.UrlShortStorage
	mxBase       sync.RWMutex
	mxShort      sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		storageBase:  make(map[string]model.UrlBaseStorage),
		storageShort: make(map[string]model.UrlShortStorage),
	}
}

func (r *Repository) GetBaseURL(url string) (model.UrlShortStorage, error) {
	r.mxBase.RLock()
	defer r.mxBase.RUnlock()
	if v, ok := r.storageShort[url]; ok {
		return v, nil
	}
	return model.UrlShortStorage{}, fmt.Errorf("short link for \"%s\" not found", url)
}

func (r *Repository) GetShortURL(url string) (model.UrlBaseStorage, error) {
	r.mxShort.RLock()
	defer r.mxShort.RUnlock()
	if v, ok := r.storageBase[url]; ok {
		return v, nil
	}
	return model.UrlBaseStorage{}, fmt.Errorf("base link for \"%s\" not found", url)

}

func (r *Repository) WriteNewLink(url, short string) (model.UrlBaseStorage, error) {
	r.mxBase.Lock()
	r.mxShort.Lock()
	defer r.mxBase.Unlock()
	defer r.mxShort.Unlock()

	r.storageBase[url] = model.UrlBaseStorage{ShortURL: short}
	r.storageShort[short] = model.UrlShortStorage{BaseURL: url}
	return r.storageBase[url], nil
}
