package memdb

import (
	"context"
	"fmt"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"go.uber.org/zap"
	"sync"
)

type Repository struct {
	storageBase  map[string]string
	storageShort map[string]string
	mxBase       sync.RWMutex
	mxShort      sync.RWMutex
}

func NewRepository() *Repository {
	logger.Debug("Creating local repository")
	return &Repository{
		storageBase:  make(map[string]string),
		storageShort: make(map[string]string),
	}
}

func (r *Repository) GetBaseURL(ctx context.Context, url string) (string, error) {
	logger.Debug("Repo:Getting base URL from local storage", zap.String("url", url))
	r.mxBase.RLock()
	defer r.mxBase.RUnlock()
	if v, ok := r.storageShort[url]; ok {
		return v, nil
	}
	logger.Warn("Couldn't find base URL", zap.String("shorturl", url))
	return "", fmt.Errorf("short link for \"%s\" not found", url)
}

func (r *Repository) GetShortURL(ctx context.Context, url string) (string, error) {
	logger.Debug("Repo:Getting short URL from local storage", zap.String("url", url))
	r.mxShort.RLock()
	defer r.mxShort.RUnlock()
	if v, ok := r.storageBase[url]; ok {
		return v, nil
	}
	logger.Warn("Couldn't find short URL", zap.String("baseurl", url))
	return "", fmt.Errorf("base link for \"%s\" not found", url)

}

func (r *Repository) WriteNewLink(ctx context.Context, url, short string) (string, error) {
	logger.Debug("Repo:Write new URL to local storage", zap.String("baseurl", url), zap.String("shorturl", short))
	if v, err := r.GetShortURL(ctx, url); err == nil {
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
