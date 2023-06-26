package memdb

import (
	"fmt"
	"sync"
)

type Repository struct {
	storageBase  map[string]string
	storageShort map[string]string
	mxBase       sync.RWMutex
	mxShort      sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		storageBase:  make(map[string]string),
		storageShort: make(map[string]string),
	}
}

func (r *Repository) GetBaseURL(url string) (string, error) {
	r.mxBase.RLock()
	defer r.mxBase.RUnlock()
	if v, ok := r.storageShort[url]; ok {
		return v, nil
	}
	return "", fmt.Errorf("short link for \"%s\" not found", url)
}

func (r *Repository) GetShortURL(url string) (string, error) {
	r.mxShort.RLock()
	defer r.mxShort.RUnlock()
	if v, ok := r.storageBase[url]; ok {
		return v, nil
	}
	return "", fmt.Errorf("base link for \"%s\" not found", url)

}

func (r *Repository) WriteNewLink(url, short string) (string, error) {
	if v, err := r.GetShortURL(url); err == nil {
		return v, nil
	}
	r.mxBase.Lock()
	r.mxShort.Lock()
	defer r.mxBase.Unlock()
	defer r.mxShort.Unlock()

	r.storageBase[url] = short
	r.storageShort[short] = url
	return r.storageBase[url], nil
}
