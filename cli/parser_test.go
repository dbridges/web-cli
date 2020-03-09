package cli

import (
	"testing"
)

func TestParse(t *testing.T) {
	type testCase struct {
		in       []string
		expected Opts
	}

	testCases := []testCase{
		{
			[]string{"web", "test"},
			Opts{Command: "open", Name: "test"},
		},
		{
			[]string{"web", "add", "test_name", "test_url"},
			Opts{Command: "add", Name: "test_name", URL: "test_url"},
		},
		{
			[]string{"web", "remove", "test_name"},
			Opts{Command: "remove", Name: "test_name"},
		},
		{
			[]string{"web", "rm", "test_name"},
			Opts{Command: "remove", Name: "test_name"},
		},
		{
			[]string{"web", "list"},
			Opts{Command: "list"},
		},
		{
			[]string{"web", "ls"},
			Opts{Command: "list"},
		},
	}

	for _, testCase := range testCases {
		opts, err := Parse(testCase.in)
		if err != nil {
			t.Fatalf("Expected not to receive error %s", err)
		}
		if opts != testCase.expected {
			t.Fatalf("Received: %v, expected: %v", opts, testCase.expected)
		}
	}
}
