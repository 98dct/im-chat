package websocket

import "net/http"

type dialOption struct {
	pattern string
	header  http.Header
}

type DialOptions func(option *dialOption)

func newDialOptions(opts ...DialOptions) dialOption {

	o := dialOption{
		pattern: "/ws",
		header:  nil,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return o
}

func WithClientPattern(pattern string) DialOptions {
	return func(option *dialOption) {
		option.pattern = pattern
	}
}

func WithClientHeader(header http.Header) DialOptions {
	return func(option *dialOption) {
		option.header = header
	}
}
