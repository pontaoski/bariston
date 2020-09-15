package templating

import (
	"net/http"
	"strings"
)

type dummyByte struct {
	sb strings.Builder
}

func (d dummyByte) Header() http.Header {
	return make(http.Header)
}

func (d dummyByte) WriteHeader(int) {}

func (d *dummyByte) Write(b []byte) (int, error) {
	return d.sb.Write(b)
}

func (d dummyByte) String() string {
	return d.sb.String()
}
