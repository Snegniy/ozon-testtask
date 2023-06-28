package service

import (
	"fmt"
	"github.com/Snegniy/ozon-testtask/internal/model"
	mock_service "github.com/Snegniy/ozon-testtask/internal/service/mocks"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	_ "github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

func TestService_GetBaseLink(t *testing.T) {
	logger.Init("yes")
	tests := map[string]struct {
		input          string
		mock           func(service *mock_service.MockRepository, shortlink string)
		expectedOutput model.UrlStorage
		expectedError  error
	}{
		"Ok": {
			input: "asdasdasdw",
			mock: func(storage *mock_service.MockRepository, shortlink string) {
				storage.EXPECT().GetBaseURL(context.Background(), shortlink).Return("https://ozon.ru", nil)
			},
			expectedOutput: model.UrlStorage{BaseURL: "https://ozon.ru"},
			expectedError:  nil,
		},
		"Not found": {
			input: "asdasdasdw",
			mock: func(storage *mock_service.MockRepository, shortlink string) {
				storage.EXPECT().GetBaseURL(context.Background(), shortlink).Return("", fmt.Errorf("base link for \"%s\" not found", shortlink))
			},
			expectedOutput: model.UrlStorage{},
			expectedError:  fmt.Errorf("base link for \"%s\" not found", "f"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_service.NewMockRepository(ctrl)

	serv := NewService(storage)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock(storage, tc.input)

			ans, err := serv.GetBaseLink(context.Background(), tc.input)
			if tc.expectedError != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedOutput, ans)
			}
		})
	}
}

/*func TestService_GetShortLink(t *testing.T) {
	logger.Init("yes")
	tests := map[string]struct {
		input          string
		mock           func(service *mock_service.MockRepository, baselink string)
		expectedOutput model.UrlStorage
		expectedError  error
	}{
		"Ok": {
			input: "ozon.ru",
			mock: func(storage *mock_service.MockRepository, baselink string) {
				storage.EXPECT().GetShortURL(context.Background(), baselink).Return("asdasdasdw", nil)
			},
			expectedOutput: model.UrlStorage{ShortURL: "asdasdasdw"},
			expectedError:  nil,
		},
		"Not found": {
			input: "ozon.ru",
			mock: func(storage *mock_service.MockRepository, baselink string) {
				storage.EXPECT().GetShortURL(context.Background(), baselink).Return("", fmt.Errorf("base link for \"%s\" not found", baselink))
			},
			expectedOutput: model.UrlStorage{},
			expectedError:  fmt.Errorf("base link for \"%s\" not found", "f"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_service.NewMockRepository(ctrl)

	serv := NewService(storage)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock(storage, tc.input)

			ans, err := serv.GetBaseLink(context.Background(), tc.input)
			if tc.expectedError != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedOutput, ans)
			}
		})
	}
}*/
