package handlers

import (
	"reflect"
	"testing"
)

func TestGridPagesAndChunk(t *testing.T) {
	cases := []struct {
		total int
		pages int
		chunk int
	}{
		{total: 1, pages: 0, chunk: 1},
		{total: 4, pages: 0, chunk: 4},
		{total: 5, pages: 2, chunk: 3},
		{total: 6, pages: 2, chunk: 3},
		{total: 9, pages: 3, chunk: 3},
		{total: 16, pages: 4, chunk: 4},
		{total: 17, pages: 4, chunk: 5},
		{total: 20, pages: 4, chunk: 5},
		{total: 36, pages: 4, chunk: 9},
		{total: 100, pages: 4, chunk: 9},
	}

	for _, c := range cases {
		if got := gridPages(c.total); got != c.pages {
			t.Errorf("gridPages(%d) = %d, want %d", c.total, got, c.pages)
		}
		if got := gridChunkSize(c.total); got != c.chunk {
			t.Errorf("gridChunkSize(%d) = %d, want %d", c.total, got, c.chunk)
		}
	}
}

func TestCarouselImageURLs(t *testing.T) {
	cases := []struct {
		total int
		want  []string
	}{
		{total: 1, want: []string{"/images/abc/1"}},
		{total: 3, want: []string{"/images/abc/1", "/images/abc/2", "/images/abc/3"}},
		{total: 5, want: []string{"/grid/abc/1", "/grid/abc/2"}},
		{total: 20, want: []string{"/grid/abc/1", "/grid/abc/2", "/grid/abc/3", "/grid/abc/4"}},
	}

	for _, c := range cases {
		if got := carouselImageURLs("abc", c.total); !reflect.DeepEqual(got, c.want) {
			t.Errorf("carouselImageURLs(abc, %d) = %v, want %v", c.total, got, c.want)
		}
	}
}
