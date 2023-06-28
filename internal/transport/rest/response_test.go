package rest

import (
	"fmt"
	"github.com/Snegniy/ozon-testtask/internal/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_writeJSON(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		data model.UrlStorage
	}
	var tests []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, writeJSON(tt.args.w, tt.args.data), fmt.Sprintf("writeJSON(%v, %v)", tt.args.w, tt.args.data))
		})
	}
}
