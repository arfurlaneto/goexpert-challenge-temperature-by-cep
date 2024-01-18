package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoundFloatTestSuite struct {
	suite.Suite
}

func TestRoundFloatTestSuite(t *testing.T) {
	suite.Run(t, new(RoundFloatTestSuite))
}

func (s *RoundFloatTestSuite) TestRoundFloat2DecimalPlaces() {

	testCases := []struct {
		input     float64
		precision uint
		expected  float64
	}{
		{input: 123.1234, precision: 1, expected: 123.1},
		{input: 987.5678, precision: 1, expected: 987.6},
		{input: 123.1234, precision: 2, expected: 123.12},
		{input: 987.5678, precision: 2, expected: 987.57},
		{input: 123.1234, precision: 3, expected: 123.123},
		{input: 987.5678, precision: 3, expected: 987.568},
		{input: 123.1234, precision: 4, expected: 123.1234},
		{input: 987.5678, precision: 4, expected: 987.5678},
	}

	for _, testCase := range testCases {
		output := RoundFloat(testCase.input, testCase.precision)
		assert.Equal(s.T(), testCase.expected, output)
	}
}
