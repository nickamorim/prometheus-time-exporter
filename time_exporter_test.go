package main

import (
	"testing"
	"time"
)

func TestDuringOfficeHours(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	tests := []struct {
		name string
		t    time.Time
		resp bool
	}{
		{
			name: "8:59 AM on a Wednesday",
			t:    time.Date(2019, 1, 2, 8, 59, 0, 0, loc),
			resp: false,
		},
		{
			name: "9:00 AM on Wednesday",
			t:    time.Date(2019, 1, 2, 9, 0, 0, 0, loc),
			resp: true,
		},
		{
			name: "10:00 AM on Wednesday",
			t:    time.Date(2019, 1, 2, 10, 0, 0, 0, loc),
			resp: true,
		},
		{
			name: "4:59 PM on Wednesday",
			t:    time.Date(2019, 1, 2, 16, 59, 0, 0, loc),
			resp: true,
		},
		{
			name: "5:00 PM on Wednesday",
			t:    time.Date(2019, 1, 2, 17, 0, 0, 0, loc),
			resp: false,
		},
		{
			name: "10:00 AM on Saturday",
			t:    time.Date(2019, 1, 5, 10, 0, 0, 0, loc),
			resp: false,
		},
		{
			name: "10:00 AM on Sunday",
			t:    time.Date(2019, 1, 6, 10, 0, 0, 0, loc),
			resp: false,
		},
	}

	for _, test := range tests {
		resp := duringOfficeHours(test.t)
		if resp != test.resp {
			t.Fatalf("\nTest Name: %v\nExpected: %v\nReceived: %v\n", test.name, test.resp, resp)
		}
	}
}

func TestDuringWakingHours(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	tests := []struct {
		name string
		t    time.Time
		resp bool
	}{
		{
			name: "7:59 AM on a Wednesday",
			t:    time.Date(2019, 1, 2, 7, 59, 0, 0, loc),
			resp: false,
		},
		{
			name: "8:00 AM on Wednesday",
			t:    time.Date(2019, 1, 2, 8, 0, 0, 0, loc),
			resp: true,
		},
		{
			name: "10:00 AM on Wednesday",
			t:    time.Date(2019, 1, 2, 10, 0, 0, 0, loc),
			resp: true,
		},
		{
			name: "9:59 PM on Wednesday",
			t:    time.Date(2019, 1, 2, 21, 59, 0, 0, loc),
			resp: true,
		},
		{
			name: "11:30 PM on Wednesday",
			t:    time.Date(2019, 1, 2, 23, 30, 0, 0, loc),
			resp: false,
		},
		{
			name: "10:00 AM on Saturday",
			t:    time.Date(2019, 1, 5, 10, 0, 0, 0, loc),
			resp: false,
		},
		{
			name: "10:00 AM on Sunday",
			t:    time.Date(2019, 1, 6, 10, 0, 0, 0, loc),
			resp: false,
		},
	}

	for _, test := range tests {
		resp := duringWakingHours(test.t)
		if resp != test.resp {
			t.Fatalf("\nTest Name: %v\nExpected: %v\nReceived: %v\n", test.name, test.resp, resp)
		}
	}
}
