package main

import (
	"errors"
	"fmt"
	"io"
)

// Resource struct is the data of the principal resource of
// this application, which is pieces of code and they origin
type Resource struct {
	Locator string
	Scheme  SchemeType
}

// NewResource function returns a Resource given a locator
func NewResource(locator string) (Resource, error) {
	var scheme SchemeType = GetSchemeTypeFrom(locator)
	var resource Resource = Resource{
		locator,
		scheme,
	}
	if scheme == UNKNOWN {
		return resource, errors.New("Unknown scheme, the locator may be a invalid one")
	}
	return resource, nil
}

func (r *Resource) String() string {
	return fmt.Sprintf("Locator: %v\nScheme: %v", r.Locator, r.Scheme)
}

// FetchData return the data which this resource points to
func (r *Resource) FetchData() (io.ReadCloser, error) {
	resolver, err := GetResolver(r.Scheme)
	var readCloser io.ReadCloser
	if err != nil {
		return readCloser, err
	}
	return resolver.FetchData(r)
}
