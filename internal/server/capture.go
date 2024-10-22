package server

import (
	"fmt"
	"io"
	"net/http"
)

type Capture struct {
	headers map[string][]string
	body    string
}

func NewCapture(r *http.Request) (Capture, error) {
	headers := make(map[string][]string)

	for name, value := range r.Header {
		headers[name] = value
	}

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return Capture{}, fmt.Errorf("error reading request body: %s", err)
	}

	return Capture{headers: headers, body: string(buf)}, nil
}

func (c Capture) Headers() map[string][]string {
	return c.headers
}

func (c Capture) Body() string {
	return c.body
}
