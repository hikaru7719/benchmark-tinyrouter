package main

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/mux"
	"github.com/hikaru7719/tinyrouter"
)

func BenchmarkTinyRouter(b *testing.B) {
	list := readEndpoint(b)
	router := setupTinyrouter(list)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/repos/owner/repo/pages/health", nil)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
	}
}

func BenchmarkGorilla(b *testing.B) {
	list := readEndpoint(b)
	router := setupGorilla(list)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/repos/owner/repo/pages/health", nil)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
	}
}

func BenchmarkChi(b *testing.B) {
	list := readEndpoint(b)
	router := setupChi(list)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/repos/owner/repo/pages/health", nil)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
	}
}

func setupTinyrouter(list []Endpoint) *tinyrouter.TinyRouter {
	router := tinyrouter.New()
	f := func(rw http.ResponseWriter, r *http.Request) {}
	for _, e := range list {
		if e.Method == "GET" {
			router.Get(e.Path, f)
		}
		if e.Method == "POST" {
			router.Post(e.Path, f)
		}
		if e.Method == "PUT" {
			router.Put(e.Path, f)
		}
		if e.Method == "PATCH" {
			router.Patch(e.Path, f)
		}
		if e.Method == "HEAD" {
			router.Head(e.Path, f)
		}
		if e.Method == "CONNECT" {
			router.Connect(e.Path, f)
		}
		if e.Method == "OPTIONS" {
			router.Options(e.Path, f)
		}
	}
	return router
}

func setupGorilla(list []Endpoint) *mux.Router {
	router := mux.NewRouter()
	f := func(rw http.ResponseWriter, r *http.Request) {}
	for _, e := range list {
		if e.Method == "GET" {
			router.HandleFunc(e.Path, f).Methods("GET")
		}
		if e.Method == "POST" {
			router.HandleFunc(e.Path, f).Methods("POST")
		}
		if e.Method == "PUT" {
			router.HandleFunc(e.Path, f).Methods("PUT")
		}
		if e.Method == "PATCH" {
			router.HandleFunc(e.Path, f).Methods("PATCH")
		}
		if e.Method == "HEAD" {
			router.HandleFunc(e.Path, f).Methods("HEAD")
		}
		if e.Method == "CONNECT" {
			router.HandleFunc(e.Path, f).Methods("CONNECT")
		}
		if e.Method == "OPTIONS" {
			router.HandleFunc(e.Path, f).Methods("OPTIONS")
		}
	}
	return router
}

func setupChi(list []Endpoint) *chi.Mux {
	router := chi.NewMux()
	f := func(rw http.ResponseWriter, r *http.Request) {}
	for _, e := range list {
		if e.Method == "GET" {
			router.Get(e.Path, f)
		}
		if e.Method == "POST" {
			router.Post(e.Path, f)
		}
		if e.Method == "PUT" {
			router.Put(e.Path, f)
		}
		if e.Method == "PATCH" {
			router.Patch(e.Path, f)
		}
		if e.Method == "HEAD" {
			router.Head(e.Path, f)
		}
		if e.Method == "CONNECT" {
			router.Connect(e.Path, f)
		}
		if e.Method == "OPTIONS" {
			router.Options(e.Path, f)
		}
	}
	return router
}

type Endpoint struct {
	Method string
	Path   string
}

func readEndpoint(b *testing.B) []Endpoint {
	b.Helper()
	f, err := os.Open("endpoint.txt")
	if err != nil {
		b.Fatal("can't read text %w", err)
	}
	scanner := bufio.NewScanner(f)
	endpointList := make([]Endpoint, 0)
	for scanner.Scan() {
		list := strings.Fields(scanner.Text())
		if len(list) != 2 {
			b.Fatal("unexpected file format")
		}
		endpointList = append(endpointList, Endpoint{Method: list[0], Path: list[1]})
	}
	return endpointList
}
