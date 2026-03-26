package repository

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUrlStorage(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func() *redis.Client
		inputCode string
		inputURL string 

		expectedErr error
		verifyFunc func(ctx context.Context, r *redis.Client,inputCode, inputURL string)
	}{
		{
			name: "normal case",

			setupMock: func() *redis.Client {
				mock := redispkg.InitMockRedis(t)
				return mock
			},

			inputCode:"1234567",
			inputURL:"https://www.google.com",
 
			
			expectedErr: nil,
			verifyFunc: func(ctx context.Context, r *redis.Client,inputCode, inputURL string) {
				res, err := r.Get(ctx, inputCode).Result()
				assert.NoError(t, err)
				assert.Equal(t, inputURL, res)

			},

		},
	}
	for _, tc := range testCases {

		t.Run(tc.name, func (t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			redisClient := tc.setupMock()
			testRepo := NewURLStorage(redisClient)

			err := testRepo.StoreURL(ctx, inputCode, inputURL)
			if err == nil {
				tc.verifyFunc(ctx, redisClient, tc.inputCode, tc.inputURL)
			}

		})
	}
}