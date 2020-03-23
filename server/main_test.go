package main

import "testing"

func TestGetYear(t *testing.T) {
	MinYear := 2020
	MaxYear := 2999
	r := GetYear()
	if r < MinYear || r > MaxYear {
		t.Errorf("GetYear: Test failed: Invalid value :%d", r)
	}
}
