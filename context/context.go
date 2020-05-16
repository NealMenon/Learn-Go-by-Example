package main

import (
	"fmt"
	"net/http"
)

// Store interface
type Store interface {
	Fetch() string
}

// StubStore is a store
type StubStore struct {
	response string
}

// Server returns required data about store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}

// Fetch returns string data
func (s *StubStore) Fetch() string {
	return ""
}
