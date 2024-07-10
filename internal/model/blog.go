package model

import "sync"

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Blog struct {
	mu       sync.RWMutex
	articles []Article
}

func (b *Blog) Add(article Article) {
	b.mu.Lock()
	b.articles = append(b.articles, article)
	b.mu.Unlock()
}

func (b *Blog) FetchAll() []Article {
	return b.articles
}
