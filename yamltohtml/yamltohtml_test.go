package yamltohtml

import "testing"

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestYamlToHtml(t *testing.T) {
	testCases := []TestCase{
		{desc: "Test Case 1", path: "testdata/test1.yml", expected: "<html><head><title>Test 1</title></head><body>Test 1 Body</body></html>"},
		{desc: "Test Case 2", path: "testdata/test2.yml", expected: "<html><head><title>Test 2</title></head><body>Test 2 Body</body></html>"},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			actual, err := YamlToHTML(test.path)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if actual != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, actual)
			}
		})
	}
}
