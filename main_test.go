package main

import (
	"errors"
	"fmt"
	"testing"
)

type SuitCase struct {
	Case         int
	ExpextResult int
	ExpextErr    error
}

func Test_Add(t *testing.T) {
	usecases := []SuitCase{
		{
			Case:         1,
			ExpextResult: 11,
			ExpextErr:    nil,
		},
		{
			Case:         2,
			ExpextResult: 12,
			ExpextErr:    nil,
		},
		{
			Case:         3,
			ExpextResult: 13,
			ExpextErr:    nil,
		},
		{
			Case:         4,
			ExpextResult: 14,
			ExpextErr:    nil,
		},
		{
			Case:         5,
			ExpextResult: 15,
			ExpextErr:    nil,
		},
		{
			Case:         -1,
			ExpextResult: -1,
			ExpextErr:    fmt.Errorf("SUBZERO_ERROR"),
		},
	}

	for i, uc := range usecases {
		actual, err := add(uc.Case)

		if actual != uc.ExpextResult {
			t.Errorf("FIF: CASE_NUMBER (%d): EXPECTED: %d, but GETTING: %d\n", i, uc.ExpextResult, actual)
		}

		if i == 5 {
			if errors.Is(err, uc.ExpextErr) {
				t.Errorf("SIF: CASE_NUMBER (%d): EXPECTED: %v, but GETTING: %v\n", i, uc.ExpextErr, err)
			}
		}
	}
}
