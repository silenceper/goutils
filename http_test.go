package goutils

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostFormWithFile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}))
	defer srv.Close()
	client := &http.Client{}
	data := make(map[string]io.Reader)
	data["field1"] = strings.NewReader("file1")
	_, err := PostFormWithFile(client, srv.URL, data)
	if err != nil {
		t.Errorf("%s", err)
	}
}
