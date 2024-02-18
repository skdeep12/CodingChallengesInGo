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
	t := domain.NewTestShortener(repository.NewInMemoryStore())
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
		re := regexp.MustCompile(`/.*`)
		matches := re.FindStringSubmatch(r.URL.Path)
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(matches)
		if len(matches) >= 1 && len(matches[0]) == 7 {
			shortUrl := matches[0][1:]
			fmt.Println(shortUrl)
			if resolvedUrl, err := t.Resolve(r.Context(), shortUrl); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusMovedPermanently)
				w.Write([]byte(resolvedUrl))
			}
		} else {
			http.NotFound(w, r)
		}
	})
	http.ListenAndServe(":9999", nil)
}
