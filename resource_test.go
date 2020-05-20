package main

import "testing"

func TestNewResourceShouldWorkAsExpected(t *testing.T) {
	locator := "lot:QmT5NvUtoM5nWFfrQdVrFtvGfKFmG7AHE8P34isapyhCxX"
	resource, err := NewResource(locator)
	if err != nil {
		t.Error(err)
	}
	if resource.Scheme != LOT {
		t.Errorf("the scheme of %s should be %s, found %s", locator, LOT, resource.Scheme)
	}
	if resource.Locator != locator {
		t.Errorf("the locator should be %s, found %s", locator, resource.Locator)
	}
}
