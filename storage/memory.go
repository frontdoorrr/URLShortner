package storage

import "sync"

var urlStore sync.Map

func SaveURL(shortURL, originalURL string) {
	urlStore.Store(shortURL, originalURL)
}

func GetURL(shortURL string) (string, bool) {
	value, ok := urlStore.Load(shortURL)
	if !ok {
		return "", false
	}
	return value.(string), true
}
