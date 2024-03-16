package shortener

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncode(t *testing.T) {
	req, err := http.NewRequest("GET", "/encode?url=http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Encode)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Short URL: aHR0cDovL2V4YW1wbGUuY29t"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDecode(t *testing.T) {
	urlMap["aHR0cDovL2V4YW1wbGUuY29t"] = "http://example.com"

	req, err := http.NewRequest("GET", "/decode/aHR0cDovL2V4YW1wbGUuY29t", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Decode)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFound)
	}

	expected := "<a href=\"http://example.com\">Found</a>.\n\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
