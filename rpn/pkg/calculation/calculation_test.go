package calculation_test

import (
	"testing"

	"github.com/asiafrolova/Calculator/rpn/pkg/calculation"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "/",
			expression:     "1/2",
			expectedResult: 0.5,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", val, testCase.expectedResult)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:        "invalid operation *",
			expression:  "1+1*",
			expectedErr: calculation.ErrInvalidExpression,
		},
		{
			name:        "invalid operation **",
			expression:  "2+2**2",
			expectedErr: calculation.ErrInvalidExpression,
		},
		{
			name:        "invalid operation (",
			expression:  "((2+2-*(2",
			expectedErr: calculation.ErrInvalidExpression,
		},
		{
			name:        "empty",
			expression:  "",
			expectedErr: calculation.ErrEmptyExp,
		},
		{
			name:        "division by zero",
			expression:  "2/0",
			expectedErr: calculation.ErrDivisionByZero,
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err != testCase.expectedErr {
				t.Fatalf("expression %s is invalid but result  %f was obtained", testCase.expression, val)
			}
		})
	}
}