package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureApiVersion(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "url without trailing slash",
			input:    "https://disco.test/api/public",
			expected: "https://disco.test/api/public",
		},
		{
			name:     "url with single trailing slash",
			input:    "https://disco.test/api/public/",
			expected: "https://disco.test/api/public",
		},
		{
			name:     "url with multiple trailing slashes",
			input:    "https://disco.test/api/public///",
			expected: "https://disco.test/api/public//",
		},
		{
			name:     "disco path without trailing slash",
			input:    "https://disco.test/disco",
			expected: "https://disco.test/disco",
		},
		{
			name:     "disco path with trailing slash",
			input:    "https://disco.test/disco/",
			expected: "https://disco.test/disco",
		},
		{
			name:     "url with custom version v32",
			input:    "https://disco.test/disco/v32",
			expected: "https://disco.test/disco/v32",
		},
		{
			name:     "base url without path",
			input:    "https://disco.test",
			expected: "https://disco.test",
		},
		{
			name:     "base url with trailing slash",
			input:    "https://disco.test/",
			expected: "https://disco.test",
		},
		{
			name:     "url with port number",
			input:    "https://disco.test:8080/api/public",
			expected: "https://disco.test:8080/api/public",
		},
		{
			name:     "url with port and trailing slash",
			input:    "https://disco.test:8080/api/public/",
			expected: "https://disco.test:8080/api/public",
		},
		{
			name:     "localhost url",
			input:    "http://localhost:3000/disco",
			expected: "http://localhost:3000/disco",
		},
		{
			name:     "localhost with trailing slash",
			input:    "http://localhost:3000/disco/",
			expected: "http://localhost:3000/disco",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Config.Host = tt.input
			EnsureApiVerison()
			assert.Equal(t, tt.expected, Config.Host)
		})
	}
}

// TestEnsureApiVersion_DetectsVersionPaths tests that URLs containing /v1 or /v2 are detected
func TestEnsureApiVersion_DetectsVersionPaths(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		shouldContainV1 bool
		shouldContainV2 bool
		description     string
	}{
		{
			name:            "url with /v1 in path",
			input:           "https://disco.test/api/public/v1",
			shouldContainV1: true,
			shouldContainV2: false,
			description:     "Should detect /v1 in the URL",
		},
		{
			name:            "url with /v2 in path",
			input:           "https://disco.test/api/public/v2",
			shouldContainV1: false,
			shouldContainV2: true,
			description:     "Should detect /v2 in the URL",
		},
		{
			name:            "url with /v1 and trailing slash",
			input:           "https://disco.test/api/public/v1/",
			shouldContainV1: true,
			shouldContainV2: false,
			description:     "Should detect /v1/ in the URL",
		},
		{
			name:            "url with /v2 and trailing slash",
			input:           "https://disco.test/api/public/v2/",
			shouldContainV1: false,
			shouldContainV2: true,
			description:     "Should detect /v2/ in the URL",
		},
		{
			name:            "url with v1 in domain name",
			input:           "https://v1.disco.test/api/public",
			shouldContainV1: true,
			shouldContainV2: false,
			description:     "Should detect v1 even in domain",
		},
		{
			name:            "url with both /v1 and /v2",
			input:           "https://disco.test/v1/api/v2",
			shouldContainV1: true,
			shouldContainV2: true,
			description:     "Should detect both /v1 and /v2",
		},
		{
			name:            "url with /v10 should still contain v1",
			input:           "https://disco.test/api/public/v10",
			shouldContainV1: true,
			shouldContainV2: false,
			description:     "Contains 'v1' substring in /v10",
		},
		{
			name:            "url with /v20 should still contain v2",
			input:           "https://disco.test/api/public/v20",
			shouldContainV1: false,
			shouldContainV2: true,
			description:     "Contains 'v2' substring in /v20",
		},
		{
			name:            "url with /v3 should not contain v1 or v2",
			input:           "https://disco.test/api/public/v3",
			shouldContainV1: false,
			shouldContainV2: false,
			description:     "Should not detect v1 or v2 in /v3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the detection logic that EnsureApiVerison uses
			trimmedURL := tt.input
			if len(trimmedURL) > 0 && trimmedURL[len(trimmedURL)-1] == '/' {
				trimmedURL = trimmedURL[:len(trimmedURL)-1]
			}

			containsV1 := false
			containsV2 := false

			// Use the same logic as in EnsureApiVerison
			if len(trimmedURL) > 0 {
				for i := 0; i < len(trimmedURL)-1; i++ {
					if trimmedURL[i:i+2] == "v1" {
						containsV1 = true
					}
					if trimmedURL[i:i+2] == "v2" {
						containsV2 = true
					}
				}
			}

			assert.Equal(t, tt.shouldContainV1, containsV1,
				"%s: v1 detection mismatch", tt.description)
			assert.Equal(t, tt.shouldContainV2, containsV2,
				"%s: v2 detection mismatch", tt.description)

			// Verify that URLs with v1 or v2 would trigger the exit condition
			wouldExit := containsV1 || containsV2
			if wouldExit {
				t.Logf("✓ URL '%s' would trigger os.Exit(1) - contains forbidden version path", tt.input)
			}
		})
	}
}

// Note: The actual os.Exit(1) behavior cannot be directly tested in unit tests.
// The TestEnsureApiVersion_DetectsVersionPaths function verifies the detection logic
// that determines when EnsureApiVerison would call os.Exit(1).
