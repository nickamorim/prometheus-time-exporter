package main

import "testing"

func TestBoolToFloat64(t *testing.T) {
	tests := []struct {
		name string
		b    bool
		f    float64
	}{
		{
			name: "True to 1",
			b:    true,
			f:    1,
		},
		{
			name: "False to 0",
			b:    false,
			f:    0,
		},
	}

	for _, test := range tests {
		resp := BoolToFloat64(test.b)
		if resp != test.f {
			t.Fatalf("\nTest Name: %v\nExpected: %v\nReceived: %v\n", test.name, test.f, resp)
		}
	}
}
