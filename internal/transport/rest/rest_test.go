package rest

import (
	"bytes"
	"github.com/Snegniy/ozon-testtask/internal/model"
	mock_rest "github.com/Snegniy/ozon-testtask/internal/transport/rest/mocks"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_PostLink(t *testing.T) {

	testTable := map[string]struct {
		body                 string
		dto                  model.UrlStorage
		mock                 func(s *mock_rest.MockServices, dto model.UrlStorage)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		"Ok": {
			body: `{"url":"ozon.ru"}`,
			dto:  model.UrlStorage{BaseURL: "ozon.ru"},
			mock: func(s *mock_rest.MockServices, dto model.UrlStorage) {
				s.EXPECT().GetShortLink(context.Background(), dto).Return(model.UrlStorage{ShortURL: "test"}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"shorturl":"test"}`,
		},
		"Bad format": {
			body:                 `{"uurrll":"ozon.ru"}`,
			dto:                  model.UrlStorage{},
			mock:                 func(s *mock_rest.MockServices, dto model.UrlStorage) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "incorrect format",
		},
		"Not found": {
			body:                 `{"url":""}`,
			dto:                  model.UrlStorage{},
			mock:                 func(s *mock_rest.MockServices, dto model.UrlStorage) {},
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: "url cannot be empty",
		},
	}

	for name, tc := range testTable {
		t.Run(name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_rest.NewMockServices(c)
			tc.mock(service, tc.dto)
			h := NewHttpHandlers(service)

			router := chi.NewRouter()

			router.Post("/", h.PostLink)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(tc.body))
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
			router.ServeHTTP(w, req)

			require.Equal(t, tc.expectedStatusCode, w.Code)
			require.Equal(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}
