package service

import (
	"regexp"
	"testing"

	"github.com/go-playground/assert/v2"
)

var urlSafeRegex = regexp.MustCompile("^[a-zA-Z0-9_-]+$")

func TestPasswordService_GeneratePassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		expectedLen int
		expectErr   error
	}{
		{
			name:        "normal case",
			expectedLen: 10,
			expectErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testSvc := NewPassWordService()

			pass, err := testSvc.GeneratePassword()

			assert.Equal(t, tc.expectedLen, len(pass))
			assert.Equal(t, tc.expectErr, err)
			assert.Equal(t, urlSafeRegex.MatchString(pass), true)
		})
	}
}
