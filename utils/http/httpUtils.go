package httputils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func StatusOK(w *httptest.ResponseRecorder) bool {
	return w.Result().StatusCode == http.StatusOK
}

func FormatError(w *httptest.ResponseRecorder) interface{} {
	var err interface{}
	json.Unmarshal(w.Body.Bytes(), &err)
	return err
}
