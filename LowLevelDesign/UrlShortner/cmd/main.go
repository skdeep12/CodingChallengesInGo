package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"shortner/domain"
	"shortner/repository"
)

type UrlShortenerRequestSchema struct {
	Url string `json:"url"`
}

type UrlShortenerResponseSchema struct {
	Key      string `json:"key"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func main() {
	s := repository.NewInMemoryStore()
	t := domain.NewTestShortener(s)
	l := domain.NewUrlLifeCycleHandler(s)
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var (
			requestSchema UrlShortenerRequestSchema
		)
		if requestBytes, err := io.ReadAll(r.Body); err != nil {
			w.Write([]byte("error in reading request body"))
			w.WriteHeader(http.StatusBadRequest)
		} else if err = json.Unmarshal(requestBytes, &requestSchema); err != nil {
			w.Write([]byte("error in unmarshilng request body"))
			w.WriteHeader(http.StatusBadRequest)
		} else if short, uErr := t.Shorten(context.Background(), requestSchema.Url); err != nil {
			http.Error(w, uErr.Error(), http.StatusExpectationFailed)
		} else {
			response := UrlShortenerResponseSchema{
				Key:      short,
				ShortUrl: fmt.Sprintf("http://localhost:9999/%s", short),
				LongUrl:  requestSchema.Url,
			}
			responseBytes, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseBytes)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("resolver")
		shortUrl := GetUrlPathForResolver(r)
		switch r.Method {
		case http.MethodGet:
			if len(shortUrl) > 0 {
				if resolvedUrl, err := t.Resolve(r.Context(), shortUrl); err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
				} else {
					w.Header().Set("Location", resolvedUrl)
					w.WriteHeader(http.StatusFound)
				}
			} else {
				http.NotFound(w, r)
			}
		case http.MethodDelete:
			if len(shortUrl) > 0 {
				if err := l.Delete(r.Context(), shortUrl); err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			} else {
				http.NotFound(w, r)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			return

		}
	})
	http.ListenAndServe(":9999", nil)
}

func GetUrlPathForResolver(r *http.Request) string {
	re := regexp.MustCompile(`/.*`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) >= 1 && len(matches[0]) == 7 {
		return matches[0][1:]
	}
	return ""
}
