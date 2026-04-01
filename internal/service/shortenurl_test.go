package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockURLStorage struct {
	storeFn func(ctx context.Context, code, url string) error
}

func (m *mockURLStorage) StoreURL(ctx context.Context, code, url string) error {
	return m.storeFn(ctx, code, url)
}

type mockCodeGen struct {
	genFn func(length int) (string, error)
}

func (m *mockCodeGen) GenerateCode(length int) (string, error) {
	return m.genFn(length)
}

func TestShortenURLService_ShortenURL(t *testing.T) {
	t.Parallel()

	test := []struct {
		name        string
		inputURL    string
		genFn       func(length int) (string, error)
		storeFn     func(ctx context.Context, code, url string) error
		expectCode  string
		expectError error
	}{
		{
			name:     "generate code failed",
			inputURL: "https://google.com",
			genFn: func(length int) (string, error) {
				return "", errors.New("gen error")
			},

			storeFn: func(ctx context.Context, code, url string) error {
				return nil
			},
			expectCode:  "",
			expectError: errors.New("gen error"),
		},
		{
			name:     "store url failed",
			inputURL: "https://google.com",
			genFn: func(length int) (string, error) {
				return "Ab12Xyz", nil
			},
			storeFn: func(ctx context.Context, code, url string) error {
				return errors.New("store error")
			},

			expectCode:  "",
			expectError: errors.New("store error"),
		},
		{
			name:     "success",
			inputURL: "https://google.com",
			genFn: func(length int) (string, error) {
				return "Ab12Xyz", nil
			},
			storeFn: func(ctx context.Context, code, url string) error {
				return nil
			},
			expectCode:  "Ab12Xyz",
			expectError: nil,
		},
	}
	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svc := NewShortenURLService(
				&mockURLStorage{storeFn: tt.storeFn},
				&mockCodeGen{genFn: tt.genFn},
			)

			code, err := svc.ShortenURL(context.Background(), tt.inputURL)

			if tt.expectError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectError.Error(), err.Error())
				assert.Equal(t, tt.expectCode, code)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectCode, code)
		})
	}
}
