package main

import (
	"errors"
	"fmt"
	"testing"
)

func Test_ValidateAnswer(t *testing.T) {
	testCases := []struct {
		answer string
		err    error
	}{
		{"1", nil},
		{"asdsd", errors.New("Invalid number")},
		{"32a", errors.New("Invalid number")},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Validate answer %v", tc.answer), func(t *testing.T) {
			result := validAnswer(tc.answer)

			if result != nil && result.Error() != tc.err.Error() {
				t.Errorf("Expected is not equal as result")
			}
		})
	}
}
