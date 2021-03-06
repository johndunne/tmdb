// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb_test

import (
	"net/url"
	"testing"

	. "github.com/tvtio/tmdb"
)

func TestPopularMovie(t *testing.T) {
	tmdb := New()
	_, err := tmdb.PopularMovie()
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestBadURLPopularMovie(t *testing.T) {
	tmdb := New()
	tmdb.BaseURL, _ = url.Parse("http://foo.bar/")
	_, err := tmdb.PopularMovie()
	if err == nil {
		t.Errorf("It should fail because the BaseURL %s, does not exist.", tmdb.BaseURL)
	}
}

func TestPopularTV(t *testing.T) {
	tmdb := New()
	_, err := tmdb.PopularTV()
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestBadURLPopularTV(t *testing.T) {
	tmdb := New()
	tmdb.BaseURL, _ = url.Parse("http://foo.bar/")
	_, err := tmdb.PopularTV()
	if err == nil {
		t.Errorf("It should fail because the BaseURL %s, does not exist.", tmdb.BaseURL)
	}
}

func TestSearchMulti(t *testing.T) {
	tmdb := New()
	_, err := tmdb.SearchMulti("matrix")
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestBadURLSearchMulti(t *testing.T) {
	tmdb := New()
	tmdb.BaseURL, _ = url.Parse("http://foo.bar/")
	_, err := tmdb.SearchMulti("matrix")
	if err == nil {
		t.Errorf("It should fail because the BaseURL %s, does not exist.", tmdb.BaseURL)
	}
}

func TestSearchMovie(t *testing.T) {
	tmdb := New()
	_, err := tmdb.SearchMovie("matrix")
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestBadURLSearchMovie(t *testing.T) {
	tmdb := New()
	tmdb.BaseURL, _ = url.Parse("http://foo.bar/")
	_, err := tmdb.SearchMovie("matrix")
	if err == nil {
		t.Errorf("It should fail because the BaseURL %s, does not exist.", tmdb.BaseURL)
	}
}
