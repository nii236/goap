package main

import (
	"net/http"
	"net/url"
)

type Deliverer struct{}

func (d *Deliverer) Do(b []byte, to *url.URL, toDo func(b []byte, u *url.URL) error) {
	_ = toDo(b, to)
}

type HttpClient struct {
	*http.Client
}
