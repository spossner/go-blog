package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlog_Add(t *testing.T) {
	b := new(Blog)
	b.Add(Article{"First article", "The body of first article"})
	b.Add(Article{"Second article", "The body of second article"})
	b.Add(Article{"Third article", "The body of third article"})

	a := b.FetchAll()

	assert.Equal(t, 3, len(a))
	assert.Equal(t, "The body of second article", a[1].Body)
	assert.Equal(t, "Third article", a[2].Title)
}
