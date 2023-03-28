package calculator

import "testing"

type TestCase struct {
	name     string
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {

	t.Run("test for all 3 digit armstrong numbers", func(t *testing.T) {
		tests := []TestCase{
			{name: "Testing value for: 153", value: 153, expected: true},
			{name: "Testing value for: 370", value: 370, expected: true},
			{name: "Testing value for: 371", value: 371, expected: true},
			{name: "Testing value for: 407", value: 407, expected: true},
		}
		for _, testCase := range tests {
			t.Run(testCase.name, func(t *testing.T) {
				testCase.actual = CalculateIsArmstrong(testCase.value)
				if testCase.actual != testCase.expected {
					t.Errorf("Expected %v, but got %v", testCase.expected, testCase.actual)
				}
			})
		}
	})

}
func TestNegativeCalculateIsArmstrong(t *testing.T) {
	t.Run("should return false for 372", func(t *testing.T) {
		testCase := TestCase{
			value:    372,
			expected: false,
		}
		testCase.actual = CalculateIsArmstrong(testCase.value)

		if testCase.actual != testCase.expected {
			t.Errorf("Expected %v, but got %v", testCase.expected, testCase.actual)
		}
	})
	t.Run("should return false for 154", func(t *testing.T) {
		testCase := TestCase{
			value:    154,
			expected: false,
		}
		testCase.actual = CalculateIsArmstrong(testCase.value)

		if testCase.actual != testCase.expected {
			t.Errorf("Expected %v, but got %v", testCase.expected, testCase.actual)
		}
	})
}

func benchmarkCalculateIsArmstrong(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateIsArmstrong(input)
	}

}

func BenchmarkCalculateIsArmstrong370(b *testing.B) {
	benchmarkCalculateIsArmstrong(370, b)
}
func BenchmarkCalculateIsArmstrong154(b *testing.B) {
	benchmarkCalculateIsArmstrong(153, b)
}
func BenchmarkCalculateIsArmstrong371(b *testing.B) {
	benchmarkCalculateIsArmstrong(371, b)
}
