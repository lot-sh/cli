package main

import "testing"

func TestGetResolverShouldWorksWhenPassedHTTPSchemeType(t *testing.T) {
	_, err := GetResolver(HTTP)
	if err != nil {
		t.Error("Error should be no returned by GetResolver when passed a HTTP SchemeType")
	}
}

func TestGetResolverShouldReturnErrorWhenPassedUNKNOWNSchemeType(t *testing.T) {
	_, err := GetResolver(UNKNOWN)
	if err == nil {
		t.Error("Error should be returned by GetResolver when passed a UNKNONW SchemeType")
	}
}
