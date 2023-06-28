package memdb

import (
	"context"
	"fmt"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"testing"
)

var baseUrl = "ozon.ru"
var shortenedUrl = "AAAAAAAAAA"
var ctx = context.Background()

func TestRepository_WriteNewLink(t *testing.T) {
	logger.Init("no")
	db := NewRepository()

	_, err := db.WriteNewLink(ctx, baseUrl, shortenedUrl)
	if err != nil {
		t.Errorf("Save error. Expected: nil, got: %s", err.Error())
	}
	base, err := db.GetBaseURL(ctx, shortenedUrl)
	if err != nil {
		t.Fatal(err)
	}
	if base != baseUrl {
		t.Errorf("Incorrect response. Expected: %s, got: %s", baseUrl, base)
	}

	_, err = db.WriteNewLink(ctx, baseUrl, shortenedUrl)
	if err != nil {
		t.Errorf("Error not handled. Expected: nil, got: %s", err)
	}
}

func TestRepository_GetBaseURL(t *testing.T) {
	logger.Init("no")
	db := NewRepository()

	base, err := db.GetBaseURL(ctx, shortenedUrl)
	ourError := fmt.Errorf("short link for \"%s\" not found", shortenedUrl)
	if err == ourError {
		t.Fatalf("Incorrect response. Expected: %s, got: %s", ourError, err)
	}

	_, err = db.WriteNewLink(ctx, baseUrl, shortenedUrl)
	if err != nil {
		t.Fatal(err)
	}

	base, err = db.GetBaseURL(ctx, shortenedUrl)
	if err != nil {
		t.Fatal(err)
	}
	if base != baseUrl {
		t.Errorf("Incorrect response. Expected: %s, got: %s", baseUrl, base)
	}
}

func TestRepository_GetShortURL(t *testing.T) {
	logger.Init("no")
	db := NewRepository()

	_, err := db.GetShortURL(ctx, baseUrl)
	if err == nil {
		t.Errorf("Incorrect response. Expected: false, got: %t", err)
	}

	_, err = db.WriteNewLink(ctx, baseUrl, shortenedUrl)
	if err != nil {
		t.Fatal(err)
	}

	short, err := db.GetShortURL(ctx, baseUrl)
	if err != nil {
		t.Errorf("Incorrect response. Expected: false, got: %t", err)
	}
	if short != shortenedUrl {
		t.Errorf("Incorrect response. Expected: %s, got: %s", shortenedUrl, short)
	}
}
