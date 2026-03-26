package repository

func TestUrlStorage(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func() *redis.Client

		expectedErr error
		verifyFunc func(ctx context.Context, r *redis.Client)
	}{
		{
			name: "normal case",

			setupMock: func() *redis.Client {
				mock := pkgredis.InitMockRedis(t)
				return mock
			},

			inputCode:"1234567",
			inputURL:"https://www.google.com",

			
			expectedErr: nil,
			verifyFunc: func(ctx context.Context, r *redis.Client inputCode, inputURL string) {
				res, err := r.Get(ctx, key:inputCode).Result()
				assert.NoError(t, err)
				assert.Equal(t, inputURL, res)

			},

		}
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
