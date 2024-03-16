package shortener

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var urlMap = make(map[string]string)

// Encode accepts long URLs as input and generates short aliases for them.
func Encode(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	_, err := url.ParseRequestURI(longURL)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortURL := base64.URLEncoding.EncodeToString([]byte(longURL))
	urlMap[shortURL] = longURL

	fmt.Fprintf(w, "Short URL: %s", shortURL)
}

// Decode accepts short aliases as input and redirects users to the original URLs.
func Decode(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/decode/")
	longURL, ok := urlMap[shortURL]
	if !ok {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
