package main

import (
	"io"
	"net/http"
)

// HTTPResolver is a HTTP/S client whose purpose is fetch data
// with a simple API
//
// • It implements the infertace Resolver
type HTTPResolver struct {
}

// FetchData is the
func (resolver *HTTPResolver) FetchData(resource *Resource) (io.ReadCloser, error) {
	res, err := http.Get(resource.Locator)
	return res.Body, err
}
