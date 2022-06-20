package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func PlayerServer(w *httptest.ResponseRecorder, f *http.Request) {
	fmt.Fprintf(w, "20")
}
